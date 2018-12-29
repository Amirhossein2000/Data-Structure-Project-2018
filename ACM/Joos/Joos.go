package main

import (
	"fmt"
)

func Swin(s, m string) bool {
	iswin := false
	xcount := 0
	for x := range s {
		if m[0] == s[x] {
			xcount = x
			for o := range m {
				if xcount == len(s) {
					xcount = 0
				}
				if m[o] == s[xcount] {
					xcount++
					iswin = true
				} else {
					iswin = false
					break
				}
			}
		}
		if iswin {
			break
		}
	}
	return iswin
}

func main() {
	var (
		s, m string
	)
	fmt.Scan(&s)
	fmt.Scan(&m)
	if Swin(s, m) {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
