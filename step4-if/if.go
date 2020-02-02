package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hi")

	var i int = 0

	if i == 2 {
		fmt.Println("i == 2")
	} else if i == 1 {
		fmt.Println("i == 1")
	} else {
		fmt.Println("i == 0")
	}
}
