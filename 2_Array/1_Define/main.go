package main

import "fmt"

var A [10]int
var B []int
var C = [10]int{1,2,3,45,34,343}
var D = []int{1,2,3,45,34,343}

func main() {

	E:=[]int{1,2,3,4}
	F:=[4]int{1,2,3,4}

	fmt.Println("A : ",A)
	fmt.Println("B : ",B)
	fmt.Println("C : ",C)
	fmt.Println("D : ",D)
	fmt.Println("E : ",E)
	fmt.Println("F : ",F)

	E=append(E,5,6,7,8)

	fmt.Println("E : ",E)

}
