package main

import "fmt"

func LineSearch(target_map []int, value int) int {

	for i,j:=range(target_map){

		if value==j {
			return i
		}

	}

	return -1

}

func main() {

	values := []int{1, 2, 3, 4, 5, 6, 7}

	if index:=LineSearch(values, 7);index!=-1 {

		fmt.Println(7,"found in",index)

	}else {

		fmt.Println(7,"Not Found")

	}

	if index:=LineSearch(values, 10);index!=-1 {

		fmt.Println(10,"found in ",index)

	}else {

		fmt.Println(10,"Not Found")

	}

}