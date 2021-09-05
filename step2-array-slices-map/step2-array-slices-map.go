package main

import (
	"fmt"
)

func main() {

	// //*******************************************************
	// //Array:
	// //
	// //BT - There are 3 ways to create an array.
	// //*******************************************************

	// //BT - Create 3 elements in an array with initialized value. No var keyword.

	// //BT - Format: type{your_init_value_her}
	// myArray := [3]int{1, 2, 3}

	// //BT - Create an array with var keyword. All elements are automatically initilized to 0.
	// var myArray2 [3]int

	// myArray2[0] = 0

	// //BT - With var keyword.
	// var myArray3 = [3]int{1, 2, 3}

	// fmt.Println(myArray[0])
	// fmt.Println(myArray2[0])
	// fmt.Println(myArray3[0])

	// fmt.Println("Array:")
	// for i, s := range myArray3 {
	// 	fmt.Printf("Index: %d equal %d\r\n", i, s)
	// }

	// //********************************************************
	// //Slice:
	// //
	// //BT - There is no number of element specify in the square
	// //     bracket.
	// //********************************************************

	// //BT - Format: type{your_init_value_her}
	// mySlices0 := []int{1, 2, 3}

	// var mySlices2 = []int{1, 2, 3}

	// var mySlices3 = make([]int, 3)

	// var mySlices4 [1]int

	// mySlices4[0] = 100

	// // fmt.Println(mySlices[0])
	// fmt.Println(mySlices2[0])
	// fmt.Println(mySlices3[0])
	// fmt.Println(mySlices4[0])

	// //BT - The range() function will return both index and the value. This is
	// //why: i,s := range mySlices
	// //i,s := range - It means both variables: i,s get := from range function.
	// for i, s := range mySlices0 {
	// 	fmt.Printf("Index: %d equal %d\r\n", i, s)
	// }

	// //********************************************************
	// //Map:
	// //********************************************************

	// //BT - Format: type{your_init_value_her}
	// myMap := map[string]int{"CA": 55112, "MN": 55434}

	// fmt.Println(myMap["CA"])

	// fmt.Println("===============================================================")

	// //BT - Do not specify the number element with slice.
	// //     You can only use the append function to assign the value to the slice.
	// //     You can do: mySlices[0] = 1 <==Can not assigned the value to slice liked this. Must use - append - function.
	// var mySlices []int
	// slice := append(mySlices, 1, 0, 3)

	// fmt.Println(slice[0])
	// fmt.Println(slice[1])
	// fmt.Println(slice[2])
	// //BT - This is how you can add new item at the end of the slice.
	// slice = append(slice, 23)
	// fmt.Println("========================")
	// fmt.Println(slice)
	// fmt.Printf("The length: %d - The capacity: %d, content: %v\n", len(slice), cap(slice), slice)

	// for i, item := range slice {
	// 	fmt.Println("Index: ", i, " Item: ", item)
	// }

	// fmt.Println("========================")
	// //BT - You can use var keyword and the = to create a variable with assigned value to it.
	// var mySlice2 = []int{23, 45, 56}

	// fmt.Println(mySlice2)

	// fmt.Println("========================")

	// //BT - You can use this := assignment to initialize a value to a variable. Go will automatically assigned that type to the variable.
	// mySlice3 := []int{4, 5}

	// fmt.Println(mySlice3)

	// fmt.Println("")
	// fmt.Println("")
	// fmt.Println("")

	// fmt.Println("========================")
	// fmt.Println("Struct: ")
	// fmt.Println("========================")
	// type MyAddress struct {
	// 	firstName string
	// 	lastName  string
	// 	email     string
	// }

	// personAddress := []MyAddress{
	// 	{firstName: "Trong", lastName: "Tran", email: "test@test.com"},
	// 	{firstName: "Jay", lastName: "Tran", email: "test@test.com"},
	// }

	// for _, item := range personAddress {
	// 	fmt.Println(item.firstName)
	// }

	//******************************************************************************
	//Slices:
	//The idea is not using var keyword. It makes more complicate.
	//Just declare a blank slice: mySlice := []int{} <==Note the empty braces.
	//******************************************************************************

	fmt.Println("==============================================")
	fmt.Println("Array")
	fmt.Println("==============================================")

	//BT - Declare a blank slices. Notes: The braces after the type int.
	mySlice := []int{}

	//BT - Later on, add new items to it.
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 2)

	for _, item := range mySlice {
		fmt.Println(item)
	}

	//BT - The second way is initialized it with values.
	mySlice2 := []int{3, 4}

	for _, item := range mySlice2 {
		fmt.Println(item)
	}

	//******************************************************************************
	//Array:
	//The idea is not using var keyword. It makes more complicate.
	//Just declare a blank slice: myArray := [3]int{} <==Note the empty braces.
	//******************************************************************************

	fmt.Println("==============================================")
	fmt.Println("Array")
	fmt.Println("==============================================")

	myArray := [3]int{}

	myArray[0] = 100
	myArray[1] = 200

	for _, item := range myArray {
		fmt.Println(item)
	}

	//*******************************************************************************
	//BT - Array in one shot.
	//*******************************************************************************

	myArray2 := [3]int{400, 500, 600}

	for _, item := range myArray2 {
		fmt.Println(item)
	}

	//******************************************************************************
	//Map:
	//The idea is not using var keyword. It makes more complicate.
	//Just declare a blank slice: myArray := [string]interface{} <==Note the empty braces.
	//******************************************************************************

	fmt.Println("==============================================")
	fmt.Println("Map")
	fmt.Println("==============================================")

	myMap := map[string]interface{}{}

	myMap["item1"] = 100
	myMap["item2"] = 200

	for key, value := range myMap {
		fmt.Printf("%s - %d\n", key, value)
	}

	//*******************************************************************************
	//BT - Map in one shot. Note: The braces.
	//*******************************************************************************

	myMap2 := map[string]interface{}{

		"item1": 1000,
		"item2": 2000,
	}

	for i, item := range myMap2 {
		fmt.Printf("%s - %d\n", i, item)
	}
}
