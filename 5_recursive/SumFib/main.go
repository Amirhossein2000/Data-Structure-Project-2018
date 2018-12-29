package main

import "fmt"

func sumfib(input int)int{
	if input==1 {
		return 1
	}else if input==0{
		return 0
	}
	return sumfib(input-1)+sumfib(input-2)
}

func main() {

	fmt.Println(sumfib(10)) //55
	fmt.Println(sumfib(20)) //6765

}

