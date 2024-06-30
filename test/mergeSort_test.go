package main

import (
	"parallel-merge-sort/mergeSort/mergeSortSequential"
	"parallel-merge-sort/mergeSort/mergeSortWithParallelMergeAndSplit"
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
