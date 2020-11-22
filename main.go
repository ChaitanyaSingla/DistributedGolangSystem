package main

import (
	"fmt"
	"reflect"
	"sort"
)

// Task : Creating an Task interface
type Task interface {
	// Contains execute defintion
	Execute() interface{}
}

// SortTask : Create structutre to store value that will be used by Task interface to execute Task
type SortTask struct {
	InputData  interface{}
	Comparator func(i, j int) bool
}

// Execute : Will sort the string array
func (m SortTask) Execute() (outputData interface{}) {
	if reflect.TypeOf(m.InputData).Kind() == reflect.Slice {
		sort.SliceStable(m.InputData, m.Comparator)
		outputData = m.InputData
	} else {
		fmt.Println("Invalid datatype")
	}
	return outputData
}

func main() {
	var sortTask Task
	inputSortValue := []int{2, 1}
	sortTask = SortTask{InputData: inputSortValue,
		Comparator: func(i, j int) bool {
			return inputSortValue[i] < inputSortValue[j]
		}}
	fmt.Println(sortTask.Execute())

	return
}
