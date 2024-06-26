package main

import (
	"fmt"
	"parallel-merge-sort/regularMergeSort"
)

func main() {
	array := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	regularMergeSort.MergeSort(array)
	fmt.Printf("%v", array)
}
