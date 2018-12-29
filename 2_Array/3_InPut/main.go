package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var MyArray [5]int64

func main() {

	scanner:=bufio.NewScanner(os.Stdin)

	fmt.Println("this is your Array now =>",MyArray)

	for i:=0 ; i<5 ;i++{

		scanner.Scan()

		input,err:=strconv.ParseInt(scanner.Text(),10,10)

		if err!=nil {
			i--

			fmt.Println("Wrong tye ! this will just accept int .")

		}else {

			MyArray[i]=input

		}

	}

	fmt.Println("this is your Array now =>",MyArray)

}
