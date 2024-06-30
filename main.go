package main

import (
	"fmt"
	"parallel-merge-sort/mergeSort/mergeSortSequential"
	"parallel-merge-sort/mergeSort/mergeSortWithParallelMergeAndSplit"
	"parallel-merge-sort/mergeSort/mergeSortWithParallelSplit"
	"runtime"
	"slices"
	"sort"
	"time"
)

func main() {
	for _, size := range []int{10000, 20000, 50000, 100000, 200000, 500000} {
		testMergeSorts(size)
	}
}

func testMergeSorts(size int) {
	runtime.GOMAXPROCS(1)
	arrayToSort := generateDecreasingArray(size)
	start := time.Now()
	mergeSortSequential.MergeSort(arrayToSort)
	sequentialMergeSortDuration := time.Since(start).Microseconds()
	fmt.Println("Finished sorting with custom sequential merge sort in", sequentialMergeSortDuration, "us")

	arrayToSortCopy := make([]int, len(arrayToSort))
	copy(arrayToSortCopy, arrayToSort)
	startSortStableFunc := time.Now()
	slices.SortStableFunc(arrayToSortCopy, func(left, right int) int {
		return left - right
	})
	sortStableFuncDuration := time.Since(startSortStableFunc).Microseconds()
	fmt.Println("Finished sorting with SortStableFunc in", sortStableFuncDuration, "us")

	arrayToSortCopy2 := make([]int, len(arrayToSort))
	copy(arrayToSortCopy2, arrayToSort)
	startSortStable := time.Now()
	sort.Stable(sort.IntSlice(arrayToSortCopy2))
	sortStableDuration := time.Since(startSortStable).Microseconds()
	fmt.Println("Finished sorting with sort.Stable in", sortStableDuration, "us")

	startParallelSplit := time.Now()
	mergeSortWithParallelSplit.MergeSort(arrayToSort)
	parallelSplitDuration := time.Since(startParallelSplit).Microseconds()
	fmt.Println("Finished sorting with parallel split in", parallelSplitDuration, "us")

	startParallelMergeAndSplit := time.Now()
	mergeSortWithParallelMergeAndSplit.MergeSort(arrayToSort)
	parallelMergeAndSplitDuration := time.Since(startParallelMergeAndSplit).Microseconds()
	fmt.Println("Finished sorting with parallel merge and split in", parallelMergeAndSplitDuration, "us")
}

func generateDecreasingArray(size int) []int {
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = size - i
	}
	return array
}
