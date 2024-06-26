package regularMergeSort

func MergeSort(array []int) {
	mergeSortWithCopy(array, 0, len(array)-1)
}

func mergeSortWithCopy(array []int, start int, end int) {
	if start < end {
		middle := (start + end) / 2
		mergeSortWithCopy(array, start, middle)
		mergeSortWithCopy(array, middle+1, end)
		mergeWithCopy(array, start, middle, end)
	}
}

func mergeWithCopy(array []int, start, mid int, end int) {
	leftSlice := make([]int, end-start+1)
	rightSlice := make([]int, end-start+1)
	copy(leftSlice, array[start:mid+1])
	copy(rightSlice, array[mid+1:end+1])
	i := 0
	j := 0
	for k := start; k <= end; k++ {
		if i == len(leftSlice) {
			array[k] = rightSlice[j]
			j++
		} else if j == len(rightSlice) {
			array[k] = leftSlice[i]
			i++
		} else if leftSlice[i] < rightSlice[j] {
			array[k] = leftSlice[i]
			i++
		} else {
			array[k] = rightSlice[j]
			j++
		}
	}
}
