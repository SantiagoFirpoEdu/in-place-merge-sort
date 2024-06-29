package main

import (
	"fmt"
	"parallel-merge-sort/regularMergeSort"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	array := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	outputArray := make([]int, len(array))
	syncChannel := make(chan struct{})
	regularMergeSort.ParallelMergeSort(array, outputArray, syncChannel)
	<-syncChannel
	fmt.Printf("%v", array)
}
