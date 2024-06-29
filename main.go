package main

import (
	"fmt"
	"sync"
)

// ParallelMerge function that merges two sorted slices using parallelism
func parallelMerge(left, right []int, result []int, wg *sync.WaitGroup) {
	defer wg.Done()
	if len(left) == 0 {
		copy(result, right)
		return
	}
	if len(right) == 0 {
		copy(result, left)
		return
	}

	// Determine the middle of the larger array
	if len(left) < len(right) {
		left, right = right, left
	}

	mid := len(left) / 2
	otherMid := binarySearch(right, left[mid])

	// Place the middle element in its correct position in the result array
	resultMid := mid + otherMid
	result[resultMid] = left[mid]

	var leftWg, rightWg sync.WaitGroup
	leftWg.Add(1)
	rightWg.Add(1)

	// Launch parallel goroutines for the left and right parts
	go parallelMerge(left[:mid], right[:otherMid], result[:resultMid], &leftWg)
	go parallelMerge(left[mid+1:], right[otherMid:], result[resultMid+1:], &rightWg)

	leftWg.Wait()
	rightWg.Wait()
}

// binarySearch function to find the position of target in a sorted array
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)
	for low < high {
		mid := (low + high) / 2
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

// MergeSort function that sorts an array using merge sort algorithm
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	result := make([]int, len(arr))
	var wg sync.WaitGroup
	wg.Add(1)
	parallelMerge(left, right, result, &wg)
	wg.Wait()

	return result
}

// ParallelMergeSort function that sorts an array using parallel merge sort
func parallelMergeSort(arr []int, wg *sync.WaitGroup) []int {
	defer wg.Done()

	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	var left, right []int
	var leftWg, rightWg sync.WaitGroup
	leftWg.Add(1)
	go func() {
		left = parallelMergeSort(arr[:mid], &leftWg)
	}()
	rightWg.Add(1)
	go func() {
		right = parallelMergeSort(arr[mid:], &rightWg)
	}()

	leftWg.Wait()
	rightWg.Wait()

	result := make([]int, len(arr))
	var mergeWg sync.WaitGroup
	mergeWg.Add(1)
	parallelMerge(left, right, result, &mergeWg)
	mergeWg.Wait()

	return result
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Original array:", arr)

	var wg sync.WaitGroup
	wg.Add(1)
	sortedArr := parallelMergeSort(arr, &wg)
	wg.Wait()

	fmt.Println("Sorted array:", sortedArr)
}
