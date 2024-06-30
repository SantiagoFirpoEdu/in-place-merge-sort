package mergeSortWithParallelMergeAndSplit


func parallelMerge(left []int, right []int, result []int, waitChannel chan<- struct{}) {
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
	parallelMerge(left[middle+1:], right[otherMiddle:], result[resultMid+1:], rightWaitChannel)

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

func MergeSort(arr []int) []int {
	resultChannel := make(chan []int, 1)
	parallelMergeSort(arr, resultChannel)
	return <-resultChannel
}

func parallelMergeSort(array []int, resultChannel chan<- []int) {
	if len(array) <= 1 {
		resultChannel <- array
		return
	}

	mid := len(array) / 2

	leftMergeChannel := make(chan []int, 1)
	go parallelMergeSort(array[:mid], leftMergeChannel)
	rightMergeChannel := make(chan []int, 1)
	parallelMergeSort(array[mid:], rightMergeChannel)

	result := make([]int, len(array))
	mergeWaitChannel := make(chan struct{}, 1)
	parallelMerge(<-leftMergeChannel, <-rightMergeChannel, result, mergeWaitChannel)
	<-mergeWaitChannel

	resultChannel <- result
}
