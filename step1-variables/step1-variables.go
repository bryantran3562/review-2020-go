package main

import (
	"fmt"
)

//BT - This means, this function can accept multiple arguments separate by a comma
//     and it can be any type. For example, int, float, bool...etc.
func printMe(a ...interface{}) {

	for _, item := range a {
		//BT - The %v is automatically convert to match whatever you pass in your
		//      params.
		fmt.Printf("This is item: %v\n", item)
	}
}

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

	fmt.Println("=================================================")

	//BT - Declare variable without specify the var keyword.

	a1 := 3

	b1 := 3.2

	s1 := "This is a string"

	fmt.Println(a1, b1, s1)

	var myInt int = 2

	var myFloat = 2.2

	var myString = "This is a string"

	//BT - You can specify the format to match what you are passing in.
	fmt.Printf("myInt: %d\n", myInt)
	fmt.Printf("myFloat: %f\n", myFloat)
	fmt.Printf("myString: %s\n", myString)

	fmt.Println("myInt: ", myInt)
	fmt.Println("myFloat: ", myFloat)
	fmt.Println("myString: ", myString)

	printMe(1, 2.3, "This is a string")

}
