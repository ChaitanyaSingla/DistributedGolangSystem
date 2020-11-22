package distributedTaskExecuter

// MergeSortedArraysTask : Create structure to store value that will be used by Task interface to execute Task
type MergeSortedArraysTask struct {
	InputData  [][]interface{}
	Comparator func(elementOne, elementTwo interface{}) bool
}

// Execute : Will sort the string array
func (m MergeSortedArraysTask) Execute() (outputData interface{}) {
	var currentSortedArray []interface{}
	for i := 0; i < len(m.InputData); i++ {
		currentSortedArray = mergeArray(currentSortedArray, m.InputData[i], m.Comparator)
	}
	outputData = currentSortedArray

	return outputData
}

// mergeArray : will merge sorted arrays
func mergeArray(firstArray []interface{}, secondArray []interface{}, comparator func(elementOne, elementTwo interface{}) bool) (sortedArray []interface{}) {
	var firstIdx, secondIdx int

	for firstIdx < len(firstArray) && secondIdx < len(secondArray) {
		if comparator(firstArray[firstIdx], secondArray[secondIdx]) {
			sortedArray = append(sortedArray, firstArray[firstIdx])
			firstIdx++
		} else {
			sortedArray = append(sortedArray, secondArray[secondIdx])
			secondIdx++
		}
	}

	for firstIdx < len(firstArray) {
		sortedArray = append(sortedArray, firstArray[firstIdx])
		firstIdx++
	}

	for secondIdx < len(secondArray) {
		sortedArray = append(sortedArray, secondArray[secondIdx])
		secondIdx++
	}

	return sortedArray
}
