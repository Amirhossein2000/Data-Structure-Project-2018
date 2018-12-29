package main

import "fmt"

func fact(input int)int{

	if input==1 {
		return 1
	}else if input==0 {
		return 0
	}

	return input*fact(input-1)

}

func main() {

	for i:=0;i<10;i++{
		fmt.Printf("%v! = %v\n",i,fact(i))
	}

}