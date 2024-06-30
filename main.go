package main

import (
	"fmt"
	"parallel-merge-sort/mergeSort/mergeSortWithParallelMergeAndSplit"
	"parallel-merge-sort/mergeSort/mergeSortWithParallelSplit"
)

func main() {
	arrayToSort := []int{903943, 38, 27, 43, 3, 9, 82, 10, 1, 2, 3, 4, 5, 6, 7, 9, 10}
	fmt.Println(arrayToSort, "original array")
	fmt.Println(mergeSortWithParallelMergeAndSplit.ParallelMergeSort(arrayToSort), "(Parallel Merge Sort with Parallel Merge and Split, sorted array)")
	fmt.Println(mergeSortWithParallelSplit.ParallelMergeSort(arrayToSort), "(Parallel Merge Sort with Parallel Split, sorted array)")
}
