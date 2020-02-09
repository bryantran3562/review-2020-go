package main

import (
	"fmt"
)

func main() {

	//*******************************************************
	//Array:
	//
	//BT - There are 3 ways to create an array.
	//*******************************************************

	//BT - Create 3 elements in an array with initialized value. No var keyword.

	//BT - Format: type{your_init_value_her}
	myArray := [3]int{1, 2, 3}

	//BT - Create an array with var keyword. All elements are automatically initilized to 0.
	var myArray2 [3]int

	myArray2[0] = 0

	//BT - With var keyword.
	var myArray3 = [3]int{1, 2, 3}

	fmt.Println(myArray[0])
	fmt.Println(myArray2[0])
	fmt.Println(myArray3[0])

	fmt.Println("Array:")
	for i, s := range myArray3 {
		fmt.Printf("Index: %d equal %d\r\n", i, s)
	}

	//********************************************************
	//Slice:
	//
	//BT - There is no number of element specify in the square
	//     bracket.
	//********************************************************

	//BT - Format: type{your_init_value_her}
	mySlices := []int{1, 2, 3}

	var mySlices2 = []int{1, 2, 3}

	var mySlices3 = make([]int, 3)

	var mySlices4 [1]int

	mySlices4[0] = 100

	// fmt.Println(mySlices[0])
	fmt.Println(mySlices2[0])
	fmt.Println(mySlices3[0])
	fmt.Println(mySlices4[0])

	for i, s := range mySlices {
		fmt.Printf("Index: %d equal %d\r\n", i, s)
	}

	//********************************************************
	//Map:
	//********************************************************

	//BT - Format: type{your_init_value_her}
	myMap := map[string]int{"CA": 55112, "MN": 55434}

	fmt.Println(myMap["CA"])

}
