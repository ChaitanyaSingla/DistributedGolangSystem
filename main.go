package main

import (
	"fmt"
	"sort"
)

var noOfSlaves = 3

func mergeArray(firstArray, secondArray []string) (sortedArray []string) {
	var firstIdx, secondIdx int

	for firstIdx < len(firstArray) && secondIdx < len(secondArray) {
		if firstArray[firstIdx] < secondArray[secondIdx] {
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

func slave(inputArray []string) (sortedArray []string) {
	sort.Strings(inputArray)
	return inputArray
}

func master(inputArray []string) (sortedArray []string) {

	n := len(inputArray)
	partionSize := n / noOfSlaves
	index := 0
	counter := 0
	sortedArrayMap := make(map[int][]string)

	for index < n {
		incrementedIndex := index + partionSize
		currentSortedArray := slave(inputArray[index:incrementedIndex])
		sortedArrayMap[counter] = currentSortedArray
		index = incrementedIndex
		counter++
	}

	for i := 0; i < len(sortedArrayMap); i++ {
		if i == 0 {
			sortedArray = append(sortedArray, sortedArrayMap[i]...)
			continue
		}
		sortedArray = mergeArray(sortedArray, sortedArrayMap[i])
	}

	return sortedArray
}

func main() {
	inputArray := []string{"b", "a", "c", "e", "d", "f", "h", "g", "i"}
	sortedArray := master(inputArray)
	fmt.Println(sortedArray)

	return
}
