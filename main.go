package main

import (
	util "DistributedGolangSystem/distibutedTaskExecuter"
	"fmt"
	"reflect"
)

// InterfaceSlice will convert interface to slice of interface
func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func main() {
	var (
		sortTask, mergeTask util.Task
		sortedArrayData     [][]interface{}
		inputSortValue      = []int{10, 8, 5, 11, 15, 12, 1, 3, 2}
		n                   = len(inputSortValue)
		noOfSlaves          = 3
		partionSize         = n / noOfSlaves
		index               = 0
	)

	for index < n {
		incrementedIndex := index + partionSize
		currentInputSortValue := inputSortValue[index:incrementedIndex]
		sortTask = util.SortTask{InputData: currentInputSortValue,
			Comparator: func(i, j int) bool {
				return currentInputSortValue[i] < currentInputSortValue[j]
			}}
		sortedArrayData = append(sortedArrayData, InterfaceSlice(sortTask.Execute()))
		index = incrementedIndex
	}

	mergeTask = util.MergeSortedArraysTask{InputData: sortedArrayData,
		Comparator: func(elementOne, elementTwo interface{}) bool {
			if reflect.TypeOf(elementOne).Kind() == reflect.Int {
				firstElement := elementOne.(int)
				secondElement := elementTwo.(int)
				return firstElement < secondElement
			}
			return false
		}}

	fmt.Println(mergeTask.Execute())

	return
}
