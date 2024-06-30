package mergeSortWithParallelMergeAndSplit

func parallelMerge(left, right []int, result []int, waitChannel chan<- struct{}) {
	defer func() {
		waitChannel <- struct{}{}
	}()

	// If there's nothing to merge on either side, just copy the other side to the result
	if len(left) == 0 {
		copy(result, right)
		return
	}
	if len(right) == 0 {
		copy(result, left)
		return
	}

	// If the left side is smaller than the right side, swap them. Always do this to ensure the left side is shorter than the right side
	if len(left) < len(right) {
		left, right = right, left
	}

	middle := len(left) / 2
	otherMiddle := binarySearch(right, left[middle])

	resultMid := middle + otherMiddle
	result[resultMid] = left[middle]

	leftWaitChannel := make(chan struct{}, 1)
	rightWaitChannel := make(chan struct{}, 1)

	go parallelMerge(left[:middle], right[:otherMiddle], result[:resultMid], leftWaitChannel)
	go parallelMerge(left[middle+1:], right[otherMiddle:], result[resultMid+1:], rightWaitChannel)

	<-leftWaitChannel
	<-rightWaitChannel
}

// Returns the index of the first element in the array that is greater than or equal to the target
func binarySearch(array []int, target int) int {
	start, end := 0, len(array)
	for start < end {
		middle := (start + end) / 2
		if array[middle] < target {
			start = middle + 1
		} else {
			end = middle
		}
	}
	return start
}

func ParallelMergeSort(arr []int) []int {
	waitChannel := make(chan struct{}, 1)
	result := parallelMergeSort(arr, waitChannel)
	<-waitChannel
	return result
}

func parallelMergeSort(array []int, waitChannel chan<- struct{}) []int {
	defer func() {
		waitChannel <- struct{}{}
	}()

	if len(array) <= 1 {
		return array
	}

	mid := len(array) / 2

	var left []int
	var right []int
	leftWaitChannel := make(chan struct{}, 1)
	go func() {
		left = parallelMergeSort(array[:mid], leftWaitChannel)
	}()
	rightWaitChannel := make(chan struct{}, 1)
	go func() {
		right = parallelMergeSort(array[mid:], rightWaitChannel)
	}()

	<-leftWaitChannel
	<-rightWaitChannel

	result := make([]int, len(array))
	mergeWaitChannel := make(chan struct{}, 1)
	parallelMerge(left, right, result, mergeWaitChannel)
	<-mergeWaitChannel

	return result
}
