package main

import "fmt"

var a string
var b int = 10
var c  = new(int)

var (
	kk int
	nn string
	rasul bool =true
)


func main() {
	//only works in func
	aa:=10

	var c  = new(int)

	a:=*c

	a++

	fmt.Println(a,aa)

}