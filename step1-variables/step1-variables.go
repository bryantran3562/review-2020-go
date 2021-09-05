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

	// var a int = 2
	// var b float32 = 2.3
	// var s string = "This is a string"
	// var b2 bool = true

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(s)
	// fmt.Println(b2)

	// fmt.Println("=================================================")

	// //BT - Declare variable without specify the var keyword.

	// a1 := 3

	// b1 := 3.2

	// s1 := "This is a string"

	// fmt.Println(a1, b1, s1)

	// var myInt int = 2

	// var myFloat = 2.2

	// var myString = "This is a string"

	// //BT - You can specify the format to match what you are passing in.
	// fmt.Printf("myInt: %d\n", myInt)
	// fmt.Printf("myFloat: %f\n", myFloat)
	// fmt.Printf("myString: %s\n", myString)

	// fmt.Println("myInt: ", myInt)
	// fmt.Println("myFloat: ", myFloat)
	// fmt.Println("myString: ", myString)

	// printMe(1, 2.3, "This is a string")

	var myInt int = 2

	var myFloat float32 = 3.2

	var myBool bool = true

	var myString string = "This is my string"

	fmt.Printf("myInt: %v, myFloat: %v, myBool: %v, myString: %v", myInt, myFloat, myBool, myString)

	//*********************************************************************************************
	//BT - You can use := to assign a value to it and Go will inferred the correct type for you.
	//*********************************************************************************************

	myInt1 := 2

	myFloat1 := 3.2

	myBool1 := true

	myString1 := "This is my string"

	fmt.Printf("myInt1: %v, myFloat1: %v, myBool1: %v, myString1: %v", myInt1, myFloat1, myBool1, myString1)

	fmt.Println("=================================================================================================")
	fmt.Println("Slices:")
	fmt.Println("=================================================================================================")

	var mySlice []int

	mySlice = append(mySlice, 1)

	fmt.Printf("This is my slice: %v\n", mySlice[0])

	mySlice2 := []int{2, 3, 4}

	for _, item := range mySlice2 {
		fmt.Println(item)
	}

	//*************************************************************************************************
	//BT - Create a blank slice and later on assign a value to it.
	//     The easiest way is declared it without the keyword VAR and later on add value to it or
	//     initiaize it in one shot as the above example.
	//*************************************************************************************************
	mySlice10 := []int{}

	mySlice10 = append(mySlice10, 1000)

	fmt.Println(mySlice10[0])

	fmt.Println("=================================================================================================")
	fmt.Println("Array:")
	fmt.Println("=================================================================================================")

	var myArray [3]int

	myArray[0] = 1
	myArray[1] = 2

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	myArray2 := [3]int{4, 5, 6}

	for i := 0; i < len(myArray2); i++ {
		fmt.Println(myArray2[i])
	}

	fmt.Println("=================================================================================================")
	fmt.Println("Map:")
	fmt.Println("=================================================================================================")

	//BT - Note: Make sure you put in the - interface{} as a type. This will take any type as a value. In this case,
	//           we are using - int - as an example below.
	//           Don't use keyword var to declare your map.
	//It is best to do all of your variable WITHOUT var keyword.
	myMap := map[string]interface{}{
		"item1": 1,
		"item2": 2,
	}

	for i, item := range myMap {
		fmt.Printf("%v - %v\n", i, item)
	}

	myMap1 := map[string]interface{}{
		"item1": 1.2,
		"item2": 2.3,
	}

	for i, item := range myMap1 {
		fmt.Printf("%v - %v\n", i, item)
	}

	//***********************************************************************************
	//BT - This is how to create a blank map and the later on to assign value to it.
	//***********************************************************************************
	myMap3 := map[string]interface{}{}

	myMap3["item1"] = 100
	myMap3["item2"] = 200

	for i, item := range myMap3 {
		fmt.Printf("%v - %v\n", i, item)
	}

}
