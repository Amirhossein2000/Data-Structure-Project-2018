package main

import (
	"fmt"
)

var MyArray [5]int64

func main() {

	fmt.Println("this is your Array now =>",MyArray)

	for i:=int64(0) ; i<5 ;i++{

		MyArray[i]=i

	}

	fmt.Println("this is your Array now =>",MyArray)

}
