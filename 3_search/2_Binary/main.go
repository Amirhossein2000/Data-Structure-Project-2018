package main

import "fmt"

func BinarySearch(target_map []int, value int) int {

	start_index := 0
	end_index := len(target_map) - 1

	for start_index <= end_index {

		median := (start_index + end_index) / 2

		if target_map[median] < value {
			start_index = median + 1
		} else {
			end_index = median - 1
		}

	}

	if start_index == len(target_map) || target_map[start_index] != value {
		return -1
	} else {
		return start_index
	}

}

func main() {

	values := []int{1, 2, 3, 4, 5, 6, 7}

	if index:=BinarySearch(values, 5);index!=-1 {

		fmt.Println(5,"found in",index)

	}else {

		fmt.Println(5,"Not Found")

	}

	if index:=BinarySearch(values, 10);index!=-1 {

		fmt.Println(10,"found in ",index)

	}else {

		fmt.Println(10,"Not Found")

	}

}