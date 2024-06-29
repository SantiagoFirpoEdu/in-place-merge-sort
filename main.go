package main

import (
	"fmt"
	"parallel-merge-sort/mergeSort"
)

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Original array:", arr)

	sortedArr := mergeSort.ParallelMergeSort(arr)

	fmt.Println("Sorted array:", sortedArr)
}
