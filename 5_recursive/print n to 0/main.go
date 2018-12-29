package main

import "fmt"

func recur(x int) {
	if x < 0 {
		return
	}
	fmt.Println(x)

	recur(x - 1)
}

func main() {
	recur(100)
}
