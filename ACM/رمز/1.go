package main

import (
	"fmt"
)

func main() {
	var nk, find int
	var strlist = []string{}
	var safebox = [][]string{}
	var str, key string
	fmt.Scan(&nk)
	fmt.Scan(&key)
	for x := 0; x < nk; x++ {
		fmt.Scan(&str)
		for h := 0; h < len(str); h++ {
			strlist = append(strlist, string(str[h]))
		}
		safebox = append(safebox, strlist)
		strlist = nil
	}
	optimal := 0
	for x := 0; x < len(key); x++ {
		for i := 0; i < len(safebox[x]); i++ {
			if string(key[x]) == safebox[x][i] {
				find = i
			}
		}
		if (len(safebox[x]) - find) < find {
			optimal += (len(safebox[x]) - find)
		} else {
			optimal += find
		}
	}
	fmt.Println(optimal)
}
