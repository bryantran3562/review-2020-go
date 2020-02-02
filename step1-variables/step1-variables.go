package main

import (
	"fmt"
)

func main() {

	//*************************************************
	//BT - These are built in type:
	//a. integer, float, string.
	//*************************************************

	//BT - Declare variable using keyword - var
	//format: var your_variable_name type

	var a int = 2
	var b float32 = 2.3
	var s string = "This is a string"
	var b2 bool = true

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(s)
	fmt.Println(b2)

	//BT - Declare variable without specify the var keyword.

	a1 := 3

	b1 := 3.2

	s1 := "This is a string"

	fmt.Println(a1, b1, s1)

}
