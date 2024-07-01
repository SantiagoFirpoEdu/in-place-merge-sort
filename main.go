package main

import (
	"fmt"
	"os"
	"parallel-merge-sort/mergeSort/mergeSortSequential"
	"parallel-merge-sort/mergeSort/mergeSortWithParallelMergeAndSplit"
	"parallel-merge-sort/mergeSort/mergeSortWithParallelSplit"
	"parallel-merge-sort/utils"
	"slices"
	"time"
)

func main() {
	for _, size := range []int{100_000_000, 200_000_000, 300_000_000, 400_000_000} {
		fmt.Println("Size:", size)
		if os.Args[0] == "sequential" {
			fmt.Println("Sequential Merge Sort")
			testSequentialMergeSorts(size)
		} else if os.Args[0] == "parallel" {
			fmt.Println("Parallel Merge Sort")
			testParallelMergeSorts(size)
		} else {
			fmt.Println("Parallel Merge Sort")
			testParallelMergeSorts(size)
			fmt.Println("Sequential Merge Sort")
			testSequentialMergeSorts(size)
		}
		fmt.Println()
	}
}

func testSequentialMergeSorts(size int) {
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
}

func testParallelMergeSorts(size int) {
	fmt.Println("Core Count:", os.Getenv("GOMAXPROCS"))
	arrayToSort := utils.GenerateDecreasingArray(size)
	startParallelSplit := time.Now()
	mergeSortWithParallelSplit.MergeSort(arrayToSort)
	parallelSplitDuration := time.Since(startParallelSplit).Seconds()
	fmt.Println("[Parallel Split] Finished in", parallelSplitDuration, "s")

	startParallelMergeAndSplit := time.Now()
	mergeSortWithParallelMergeAndSplit.MergeSort(arrayToSort)
	parallelMergeAndSplitDuration := time.Since(startParallelMergeAndSplit).Seconds()
	fmt.Println("[Parallel Merge and Split] Finished in", parallelMergeAndSplitDuration, "s")
}
