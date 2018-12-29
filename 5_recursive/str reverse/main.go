package main

import "fmt"

func reverse(str string)string  {
	if len(str)<2 {
		return str
	}
	return string(str[len(str)-1])+reverse(string(str[:len(str)-1]))
}

func main() {

	fmt.Println(reverse("Hello"))
	fmt.Println(reverse("Amirhossein"))

}
