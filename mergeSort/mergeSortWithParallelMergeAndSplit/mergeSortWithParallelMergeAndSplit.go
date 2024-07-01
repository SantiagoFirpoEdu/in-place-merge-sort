package mergeSortWithParallelMergeAndSplit

import (
	"sync"
)

const threshold = 1000 //Threshold used to decide when to switch to sequential merge sort, adjust this value based on experimentation

func sequentialMerge(left []int, right []int, result []int) {
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}
}

func parallelMerge(left []int, right []int, result []int, wg *sync.WaitGroup) {
	defer wg.Done()

	if len(left) == 0 {
		copy(result, right)
		return
	}
	if len(right) == 0 {
		copy(result, left)
		return
	}

	if len(left) > len(right) {
		left, right = right, left
	}

	middle := len(left) / 2
	otherMiddle := binarySearch(right, left[middle])

	resultMid := middle + otherMiddle
	result[resultMid] = left[middle]

	var leftWaitGroup sync.WaitGroup
	var rightWaitGroup sync.WaitGroup

	leftWaitGroup.Add(1)
	go parallelMerge(left[:middle], right[:otherMiddle], result[:resultMid], &leftWaitGroup)

	rightWaitGroup.Add(1)
	go parallelMerge(left[middle+1:], right[otherMiddle:], result[resultMid+1:], &rightWaitGroup)

	leftWaitGroup.Wait()
	rightWaitGroup.Wait()
}

// Returns the index of the first element in the array that is greater than or equal to the target
func binarySearch(array []int, target int) int {
	start, end := 0, len(array)
	for start < end {
		middle := (start + end) / 2
		if array[middle] < target {
			start = middle + 1
		} else {
			end = middle
		}
	}
	return start
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go parallelMergeSort(arr, result, wg)

	wg.Wait()
	return result
}

func parallelMergeSort(array []int, result []int, wg *sync.WaitGroup) {
	defer wg.Done()

	if len(array) <= threshold {
		sequentialMergeSort(array, result)
		return
	}

	middle := len(array) / 2

	left := array[:middle]
	right := array[middle:]

	leftResult := result[:middle]
	rightResult := result[middle:]

	var leftWaitGroup sync.WaitGroup
	var rightWaitGroup sync.WaitGroup

	leftWaitGroup.Add(1)
	go parallelMergeSort(left, leftResult, &leftWaitGroup)

	rightWaitGroup.Add(1)
	parallelMergeSort(right, rightResult, &rightWaitGroup)

	leftWaitGroup.Wait()
	rightWaitGroup.Wait()

	sequentialMerge(leftResult, rightResult, result)
}

func sequentialMergeSort(array []int, result []int) {
	if len(array) <= 1 {
		copy(result, array)
		return
	}

	middle := len(array) / 2
	left := make([]int, middle)
	right := make([]int, len(array)-middle)

	sequentialMergeSort(array[:middle], left)
	sequentialMergeSort(array[middle:], right)
	sequentialMerge(left, right, result)
}
