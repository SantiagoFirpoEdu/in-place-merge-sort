package mergeSortWithParallelSplit

func ParallelMergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	middle := len(array) / 2
	left := ParallelMergeSort(array[:middle])
	right := ParallelMergeSort(array[middle:])

	result := make([]int, len(array))
	merge(left, right, result)

	return result
}

func merge(left []int, right []int, result []int) {
	i := 0
	j := 0
	k := 0

	// Merge the two slices while both have elements
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	// Consume remaining elements of left slice, if any
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	// Consume remaining elements of right slice, if any
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}
}
