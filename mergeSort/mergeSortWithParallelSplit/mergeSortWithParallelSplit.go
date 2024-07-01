package mergeSortWithParallelSplit

import (
	"runtime"
	"sync"
)

const threshold = 1000 // Threshold used to decide when to switch to sequential merge sort, adjust this value based on experimentation

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	outChannel := make(chan []int, 1)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go parallelMergeSort(array, outChannel, wg)
	wg.Wait()

	return <-outChannel
}

func parallelMergeSort(array []int, outChannel chan<- []int, wg *sync.WaitGroup) {
	defer wg.Done()

	if len(array) <= threshold {
		result := make([]int, len(array))
		sequentialMergeSort(array, result)
		outChannel <- result
		return
	}

	middle := len(array) / 2

	leftChannel := make(chan []int, 1)
	rightChannel := make(chan []int, 1)

	var leftWaitGroup sync.WaitGroup
	var rightWaitGroup sync.WaitGroup

	leftWaitGroup.Add(1)
	go parallelMergeSort(array[:middle], leftChannel, &leftWaitGroup)

	rightWaitGroup.Add(1)
	go parallelMergeSort(array[middle:], rightChannel, &rightWaitGroup)

	leftWaitGroup.Wait()
	rightWaitGroup.Wait()

	result := make([]int, len(array))
	sequentialMerge(<-leftChannel, <-rightChannel, result)

	outChannel <- result
}

func sequentialMergeSort(array []int, result []int) {
	if len(array) <= 1 {
		copy(result, array)
		return
	}

	mid := len(array) / 2
	left := make([]int, mid)
	right := make([]int, len(array)-mid)

	sequentialMergeSort(array[:mid], left)
	sequentialMergeSort(array[mid:], right)
	sequentialMerge(left, right, result)
}

func sequentialMerge(left []int, right []int, result []int) {
	leftPointer := 0
	rightPointer := 0
	resultPointer := 0

	for leftPointer < len(left) && rightPointer < len(right) {
		if left[leftPointer] <= right[rightPointer] {
			result[resultPointer] = left[leftPointer]
			leftPointer++
		} else {
			result[resultPointer] = right[rightPointer]
			rightPointer++
		}
		resultPointer++
	}

	for leftPointer < len(left) {
		result[resultPointer] = left[leftPointer]
		leftPointer++
		resultPointer++
	}

	for rightPointer < len(right) {
		result[resultPointer] = right[rightPointer]
		rightPointer++
		resultPointer++
	}
}
