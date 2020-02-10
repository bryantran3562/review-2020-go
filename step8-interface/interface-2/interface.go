package main

import (
	"fmt"
)

func test(i interface{}) {
	fmt.Printf("The Type is: %T - value is: %v", i, i)
	fmt.Println()
}

func main() {

	//BT - interface{} - Empty interface
	//All types are implemented interface{}

	i := 5
	s := "This is a string"
	b := true

	test(i)
	test(s)
	test(b)

}
