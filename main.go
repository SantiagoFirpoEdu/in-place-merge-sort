package main

import (
	"fmt"
	"parallel-merge-sort/mergeSort/mergeSortSequential"
	"parallel-merge-sort/mergeSort/mergeSortWithParallelMergeAndSplit"
	"parallel-merge-sort/mergeSort/mergeSortWithParallelSplit"
	"parallel-merge-sort/utils"
	"slices"
	"sort"
	"time"
)

func main() {
	for _, size := range []int{100000, 200000, 500000, 1000000, 2000000, 5000000, 10000000, 20000000, 50000000} {
		testMergeSorts(size)
	}
}

func testMergeSorts(size int) {
	arrayToSort := utils.GenerateDecreasingArray(size)
	start := time.Now()
	mergeSortSequential.MergeSort(arrayToSort)
	sequentialMergeSortDuration := time.Since(start).Milliseconds()
	fmt.Println("Finished sorting with custom sequential merge sort in", sequentialMergeSortDuration, "ms")

	arrayToSortCopy := make([]int, len(arrayToSort))
	copy(arrayToSortCopy, arrayToSort)
	startSortStableFunc := time.Now()
	slices.SortStableFunc(arrayToSortCopy, func(left, right int) int {
		return left - right
	})
	sortStableFuncDuration := time.Since(startSortStableFunc).Milliseconds()
	fmt.Println("Finished sorting with SortStableFunc in", sortStableFuncDuration, "ms")

	arrayToSortCopy2 := make([]int, len(arrayToSort))
	copy(arrayToSortCopy2, arrayToSort)
	startSortStable := time.Now()
	sort.Stable(sort.IntSlice(arrayToSortCopy2))
	sortStableDuration := time.Since(startSortStable).Milliseconds()
	fmt.Println("Finished sorting with sort.Stable in", sortStableDuration, "ms")

	startParallelSplit := time.Now()
	mergeSortWithParallelSplit.MergeSort(arrayToSort)
	parallelSplitDuration := time.Since(startParallelSplit).Milliseconds()
	fmt.Println("Finished sorting with parallel split in", parallelSplitDuration, "ms")

	startParallelMergeAndSplit := time.Now()
	mergeSortWithParallelMergeAndSplit.MergeSort(arrayToSort)
	parallelMergeAndSplitDuration := time.Since(startParallelMergeAndSplit).Milliseconds()
	fmt.Println("Finished sorting with parallel merge and split in", parallelMergeAndSplitDuration, "ms")
}
