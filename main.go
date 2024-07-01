package main

import (
	"fmt"
	"parallel-merge-sort/mergeSort/mergeSortSequential"
	"parallel-merge-sort/mergeSort/mergeSortWithParallelMergeAndSplit"
	"parallel-merge-sort/mergeSort/mergeSortWithParallelSplit"
	"parallel-merge-sort/utils"
	"runtime"
	"slices"
	"sort"
	"time"
)

func main() {
	n := 6
	runtime.GOMAXPROCS(n)
	fmt.Println("Number of CPUs used:", n)
	for _, size := range []int{100_000_000, 200_000_000, 400_000_000, 800_000_000} {
		fmt.Println("Size:", size)
		testMergeSorts(size)
		fmt.Println()
	}
}

func testMergeSorts(size int) {
	arrayToSort := utils.GenerateDecreasingArray(size)
	start := time.Now()
	mergeSortSequential.MergeSort(arrayToSort)
	sequentialMergeSortDuration := time.Since(start).Seconds()
	fmt.Println("[Sequential] Finished in", sequentialMergeSortDuration, "s")

	arrayToSortCopy := make([]int, len(arrayToSort))
	copy(arrayToSortCopy, arrayToSort)
	startSortStableFunc := time.Now()
	slices.SortStableFunc(arrayToSortCopy, func(left, right int) int {
		return left - right
	})
	sortStableFuncDuration := time.Since(startSortStableFunc).Seconds()
	fmt.Println("[SortStableFunc (Standard Library)] Finished in", sortStableFuncDuration, "s")

	arrayToSortCopy2 := make([]int, len(arrayToSort))
	copy(arrayToSortCopy2, arrayToSort)
	startSortStable := time.Now()
	sort.Stable(sort.IntSlice(arrayToSortCopy2))
	sortStableDuration := time.Since(startSortStable).Seconds()
	fmt.Println("[sort.Stable (Standard Library)] Finished in", sortStableDuration, "s")

	startParallelSplit := time.Now()
	mergeSortWithParallelSplit.MergeSort(arrayToSort)
	parallelSplitDuration := time.Since(startParallelSplit).Seconds()
	fmt.Println("[Parallel Split] Finished in", parallelSplitDuration, "s")

	startParallelMergeAndSplit := time.Now()
	mergeSortWithParallelMergeAndSplit.MergeSort(arrayToSort)
	parallelMergeAndSplitDuration := time.Since(startParallelMergeAndSplit).Seconds()
	fmt.Println("[Parallel Merge and Split] Finished in", parallelMergeAndSplitDuration, "s")
}
