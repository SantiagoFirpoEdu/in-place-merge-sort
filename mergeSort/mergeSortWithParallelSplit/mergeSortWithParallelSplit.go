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
	leftCurrentIndex := 0
	rightCurrentIndex := 0
	currentIndexInResult := 0

	// Merge the two slices while both have elements
	for leftCurrentIndex < len(left) && rightCurrentIndex < len(right) {
		if left[leftCurrentIndex] <= right[rightCurrentIndex] {
			result[currentIndexInResult] = left[leftCurrentIndex]
			leftCurrentIndex++
		} else {
			result[currentIndexInResult] = right[rightCurrentIndex]
			rightCurrentIndex++
		}
		currentIndexInResult++
	}

	//Only one of the following loops will run
	// Consume remaining elements of left slice, if any
	for leftCurrentIndex < len(left) {
		result[currentIndexInResult] = left[leftCurrentIndex]
		leftCurrentIndex++
		currentIndexInResult++
	}

	// Consume remaining elements of right slice, if any
	for rightCurrentIndex < len(right) {
		result[currentIndexInResult] = right[rightCurrentIndex]
		rightCurrentIndex++
		currentIndexInResult++
	}
}
