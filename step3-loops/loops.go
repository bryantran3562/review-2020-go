package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hi")

	//BT - Create an array

	myArray := [3]int{1, 2, 3}

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

}
