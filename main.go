package main

import (
	"fmt"
	"parallel-merge-sort/mergeSort/mergeSortSequential"
	"runtime"
)

func main() {
	n := (runtime.NumCPU() / 2) - 1
	runtime.GOMAXPROCS(n)
	arrayToSort := generateDecreasingArray(100_000)
	mergeSortSequential.MergeSort(arrayToSort)
	fmt.Println("Finished sorting")
}

func generateDecreasingArray(size int) []int {
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = size - i
	}
	return array
}
