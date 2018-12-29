package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MyArray [15]int

func main() {

	rand.Seed(time.Now().UnixNano())

	fmt.Println("this is your Array now =>",MyArray)

	for i:=0 ; i<15 ;i++{

		MyArray[i]=rand.Intn(100)

	}

	fmt.Println("this is your Array now =>",MyArray)

}
