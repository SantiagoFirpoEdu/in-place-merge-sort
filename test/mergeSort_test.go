package main

import (
	"parallel-merge-sort/mergeSort/mergeSortSequential"
	"parallel-merge-sort/mergeSort/mergeSortWithParallelMergeAndSplit"
	"parallel-merge-sort/mergeSort/mergeSortWithParallelSplit"
	"parallel-merge-sort/utils"
	"sort"
	"testing"
)

// Make some benchmarks for comparison (sequential and parallel)
func BenchmarkMergeSortWithParallelMergeAndSplit(b *testing.B) {
	array := []int{38, 27, 43, 3, 9, 82, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mergeSortWithParallelMergeAndSplit.MergeSort(array)
	}
}

func BenchmarkMergeSortWithParallelSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.ResetTimer()
		array := []int{38, 27, 43, 3, 9, 82, 10}
		mergeSortSequential.MergeSort(array)
	}
}

func BenchmarkSequentialMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.ResetTimer()
		array := []int{38, 27, 43, 3, 9, 82, 10}
		mergeSortSequential.MergeSort(array)
	}
}

func TestSequentialMergeSort(t *testing.T) {
	size := 100
	arrayToSort := utils.GenerateDecreasingArray(size)
	sequentialResult := mergeSortSequential.MergeSort(arrayToSort)
	if !sort.IsSorted(sort.IntSlice(sequentialResult)) {
		t.Errorf("Sequential merge sort is not sorted")
	}

}

func TestParallelMergeAndSplitMergeSort(t *testing.T) {
	size := 100
	arrayToSort := utils.GenerateDecreasingArray(size)
	parallelResult := mergeSortWithParallelMergeAndSplit.MergeSort(arrayToSort)
	if !sort.IsSorted(sort.IntSlice(parallelResult)) {
		t.Errorf("Parallel merge and split merge sort is not sorted")
	}
}

func TestParallelSplitMergeSort(t *testing.T) {
	size := 100
	arrayToSort := utils.GenerateDecreasingArray(size)
	parallelResult := mergeSortWithParallelSplit.MergeSort(arrayToSort)
	if !sort.IsSorted(sort.IntSlice(parallelResult)) {
		t.Errorf("Parallel split merge sort is not sorted")
	}
}
