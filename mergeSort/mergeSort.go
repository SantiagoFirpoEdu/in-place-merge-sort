package mergeSort

import (
	"sync"
)

func parallelMerge(left, right []int, result []int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	if len(left) == 0 {
		copy(result, right)
		return
	}
	if len(right) == 0 {
		copy(result, left)
		return
	}

	if len(left) < len(right) {
		left, right = right, left
	}

	middle := len(left) / 2
	otherMiddle := binarySearch(right, left[middle])

	resultMid := middle + otherMiddle
	result[resultMid] = left[middle]

	var leftWg sync.WaitGroup
	var rightWg sync.WaitGroup
	leftWg.Add(1)
	rightWg.Add(1)

	go parallelMerge(left[:middle], right[:otherMiddle], result[:resultMid], &leftWg)
	go parallelMerge(left[middle+1:], right[otherMiddle:], result[resultMid+1:], &rightWg)

	leftWg.Wait()
	rightWg.Wait()
}

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

func SequentialMergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	middle := len(array) / 2
	left := SequentialMergeSort(array[:middle])
	right := SequentialMergeSort(array[middle:])

	result := make([]int, len(array))
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	parallelMerge(left, right, result, &waitGroup)
	waitGroup.Wait()

	return result
}
func ParallelMergeSort(arr []int) []int {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	result := parallelMergeSort(arr, &waitGroup)
	waitGroup.Wait()
	return result
}

func parallelMergeSort(array []int, waitGroup *sync.WaitGroup) []int {
	defer waitGroup.Done()

	if len(array) <= 1 {
		return array
	}

	mid := len(array) / 2

	var left, right []int
	var leftWg, rightWg sync.WaitGroup
	leftWg.Add(1)
	go func() {
		left = parallelMergeSort(array[:mid], &leftWg)
	}()
	rightWg.Add(1)
	go func() {
		right = parallelMergeSort(array[mid:], &rightWg)
	}()

	leftWg.Wait()
	rightWg.Wait()

	result := make([]int, len(array))
	var mergeWg sync.WaitGroup
	mergeWg.Add(1)
	parallelMerge(left, right, result, &mergeWg)
	mergeWg.Wait()

	return result
}
