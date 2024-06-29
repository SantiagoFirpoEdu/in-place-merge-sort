package regularMergeSort

/**
 * A: Input array
 * B: Output array
 * lo: lower bound
 * hi: upper bound
 * off: offset
 */
func ParallelMergeSort(inputArray []int, outputArray []int, syncChannel chan struct{}) {
	length := uint(len(inputArray))
	if length == 1 {
		outputArray[0] = inputArray[0]
	} else {
		T := make([]int, length)
		mid := length / 2
		otherMid := mid - 1
		mySyncChannel := make(chan struct{})
		go ParallelMergeSort(inputArray[0:mid+1], T, mySyncChannel)
		go ParallelMergeSort(inputArray[mid+1:length], T[otherMid+1:length], mySyncChannel)
		<-mySyncChannel
		<-mySyncChannel
		parallelMerge(T[1:otherMid+1], T[otherMid+1:length], outputArray, syncChannel)
	}
}

func parallelMerge(left []int, right []int, outputArray []int, syncChannel chan struct{}) {

	leftLength := uint(len(left))
	rightLength := uint(len(right))
	outputLength := uint(len(outputArray))

	if leftLength < rightLength {
		//swap A and B  // ensure that A is the larger array: i, j still belong to A; k, ℓ to B

		//swap leftLength and rightLength
	}

	if leftLength <= 0 {
		return // base case, nothing to merge
	}

	leftMiddle := leftLength / 2
	s := binarySearch(left[leftMiddle], right)
	if s == -1 {
		return
	}
	t := (leftMiddle) + (uint(s))
	outputArray[t] = left[leftMiddle]

	mySyncChannel := make(chan struct{})
	go parallelMerge(left[0:leftMiddle+1], right[0:s+1], outputArray[0:t+1], mySyncChannel)
	go parallelMerge(left[leftMiddle+1:leftLength], right[s:rightLength], outputArray[t+1:outputLength], mySyncChannel)
	<-mySyncChannel
	<-mySyncChannel
	syncChannel <- struct{}{}
}

func binarySearch(key int, array []int) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2
		if array[mid] < key {
			low = mid + 1
		} else if array[mid] > key {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
