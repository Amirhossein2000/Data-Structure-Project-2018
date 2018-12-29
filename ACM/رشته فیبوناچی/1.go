package main

import (
	"fmt"
)

func main() {
	isfib := false
	var l = []int{}
	var n int
	a := 0
	b := 1
	c := 1
	fmt.Scanln(&n)
	i := 0
	for c <= n {
		i++
		l = append(l, c)
		a = b
		b = c
		c = a + b
	}
	for j := 1; j <= n; j++ {
		for x := 0; x <= (len(l) - 1); x++ {
			if j == l[x] {
				isfib = true
				break
			}
		}
		if isfib == true {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
		isfib = false
	}
}
