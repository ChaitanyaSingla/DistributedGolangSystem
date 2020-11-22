package main

import (
	util "DistributedGolangSystem/distibutedTaskExecuter"
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"reflect"
)

// Response : Response structure
type Response struct {
	Data interface{}
}

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
	createMaster := flag.Bool("createMaster", false, "master node")
	port := flag.String("port", "8000", "default port is 8000")
	masterIPAddress := flag.String("masterIPAddress", "localhost:8000", "default master ip address is 8000")
	flag.Parse()

	if *createMaster {
		listenOnPort(*port)
	} else {
		createClient(*masterIPAddress)
	}

	return
}

// listenOnPort : will create master node and will listen for other connections
func listenOnPort(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err == nil {
		fmt.Println("Master node is created and listening")
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			if err != nil {
				log.Println(err)
				continue
			}
		} else {
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

			resp := Response{Data: mergeTask.Execute()}
			b, _ := json.Marshal(resp)
			conn.Write(b)

			conn.Close()
		}
	}
}

// createClient : will create clients that will connect to master node
func createClient(masterIPAddress string) {
	conn, err := net.Dial("tcp", masterIPAddress)
	if err != nil {
		log.Fatalln(err)
	}

	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Sorted data->: " + message)
	defer conn.Close()
}
