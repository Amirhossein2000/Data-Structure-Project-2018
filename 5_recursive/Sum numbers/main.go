package main

import "fmt"

func sum(number int)int{

	if number == 1 {

		return number
	}

	return number + sum(number-1)
}

func main() {

	fmt.Println(sum(10))

}
