package main

import (
	"fmt"
)

func main() {

	i := 2

	switch i {
	case 1:
		fmt.Println("i == 1")
	case 2:
		fmt.Println("i == 2")
	default:
		fmt.Println("i == 10")
	}
}
