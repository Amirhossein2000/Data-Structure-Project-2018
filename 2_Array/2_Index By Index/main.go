package main

import "fmt"

var (
	a []int
	b [10]int
)

func main() {

	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

	a=append(a,10)
	b[0]=2
	b[7]=8
	b[3]=11

	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

}
