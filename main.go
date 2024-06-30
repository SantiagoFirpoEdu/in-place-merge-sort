package main

import (
	"fmt"
	"parallel-merge-sort/mergeSort/mergeSortWithParallelMergeAndSplit"
)

func main() {
	arrayToSort := generateDecreasingArray(10000000)
	mergeSortWithParallelMergeAndSplit.ParallelMergeSort(arrayToSort)
	fmt.Println("Finished sorting")
}

func generateDecreasingArray(size int) []int {
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = size - i
	}
	return array
}
