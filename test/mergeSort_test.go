package main

import (
	"parallel-merge-sort/mergeSort/mergeSortParallelInPlace"
	"parallel-merge-sort/mergeSort/mergeSortSequential"
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
		mergeSortParallelInPlace.MergeSort(array)
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
	size1 := 10
	arrayToSort1 := utils.GenerateDecreasingArray(size1)
	sequentialResult1 := mergeSortSequential.MergeSort(arrayToSort1)

	size2 := 100
	arrayToSort2 := utils.GenerateDecreasingArray(size2)
	sequentialResult2 := mergeSortSequential.MergeSort(arrayToSort2)

	size3 := 300
	arrayToSort3 := utils.GenerateDecreasingArray(size3)
	sequentialResult3 := mergeSortSequential.MergeSort(arrayToSort3)

	size4 := 1000
	arrayToSort4 := utils.GenerateDecreasingArray(size4)
	sequentialResult4 := mergeSortSequential.MergeSort(arrayToSort4)

	if !sort.IsSorted(sort.IntSlice(sequentialResult1)) {
		t.Errorf("Sequential merge sort is not sorted")
	}

	if !sort.IsSorted(sort.IntSlice(sequentialResult2)) {
		t.Errorf("Sequential merge sort is not sorted")
	}

	if !sort.IsSorted(sort.IntSlice(sequentialResult3)) {
		t.Errorf("Sequential merge sort is not sorted")
	}

	if !sort.IsSorted(sort.IntSlice(sequentialResult4)) {
		t.Errorf("Sequential merge sort is not sorted")
	}
}

func TestParallelMergeAndSplitMergeSort(t *testing.T) {
	size1 := 10
	arrayToSort1 := utils.GenerateDecreasingArray(size1)
	parallelResult1 := mergeSortParallelInPlace.MergeSort(arrayToSort1)

	size2 := 100
	arrayToSort2 := utils.GenerateDecreasingArray(size2)
	parallelResult2 := mergeSortParallelInPlace.MergeSort(arrayToSort2)

	size3 := 300
	arrayToSort3 := utils.GenerateDecreasingArray(size3)
	parallelResult3 := mergeSortParallelInPlace.MergeSort(arrayToSort3)

	size4 := 1000
	arrayToSort4 := utils.GenerateDecreasingArray(size4)
	parallelResult4 := mergeSortParallelInPlace.MergeSort(arrayToSort4)

	if !sort.IsSorted(sort.IntSlice(parallelResult1)) {
		t.Errorf("Parallel merge and split merge sort is not sorted")
	}

	if !sort.IsSorted(sort.IntSlice(parallelResult2)) {
		t.Errorf("Parallel merge and split merge sort is not sorted")
	}

	if !sort.IsSorted(sort.IntSlice(parallelResult3)) {
		t.Errorf("Parallel merge and split merge sort is not sorted")
	}

	if !sort.IsSorted(sort.IntSlice(parallelResult4)) {
		t.Errorf("Parallel merge and split merge sort is not sorted")
	}
}

func TestParallelSplitMergeSort(t *testing.T) {
	size1 := 10
	arrayToSort1 := utils.GenerateDecreasingArray(size1)
	parallelResult1 := mergeSortWithParallelSplit.MergeSort(arrayToSort1)

	size2 := 100
	arrayToSort2 := utils.GenerateDecreasingArray(size2)
	parallelResult2 := mergeSortWithParallelSplit.MergeSort(arrayToSort2)

	size3 := 300
	arrayToSort3 := utils.GenerateDecreasingArray(size3)
	parallelResult3 := mergeSortWithParallelSplit.MergeSort(arrayToSort3)

	size4 := 1000
	arrayToSort4 := utils.GenerateDecreasingArray(size4)
	parallelResult4 := mergeSortWithParallelSplit.MergeSort(arrayToSort4)

	if !sort.IsSorted(sort.IntSlice(parallelResult1)) {
		t.Errorf("Parallel split merge sort is not sorted")
	}

	if !sort.IsSorted(sort.IntSlice(parallelResult2)) {
		t.Errorf("Parallel split merge sort is not sorted")
	}

	if !sort.IsSorted(sort.IntSlice(parallelResult3)) {
		t.Errorf("Parallel split merge sort is not sorted")
	}

	if !sort.IsSorted(sort.IntSlice(parallelResult4)) {
		t.Errorf("Parallel split merge sort is not sorted")
	}
}
