package mergeSortWithParallelSplit

func MergeSort(array []int) []int {
	outChannel := make(chan []int, 1)
	parallelMergeSort(array, outChannel)
	return <-outChannel
}

func parallelMergeSort(array []int, outChannel chan<- []int) {
	if len(array) <= 1 {
		outChannel <- array
		return
	}

	middle := len(array) / 2
	leftChannel := make(chan []int, 1)
	go parallelMergeSort(array[:middle], leftChannel)

	rightChannel := make(chan []int, 1)
	parallelMergeSort(array[middle:], rightChannel)

	result := make([]int, len(array))
	merge(<-leftChannel, <-rightChannel, result)

	outChannel <- result
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
