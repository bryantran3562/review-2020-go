package main

import (
	"fmt"
)

//BT - Receiving 2 args. Return int
func add(a int, b int) int {
	return a + b
}

//BT - Receiving 3 args with the same type. Return int
func addMultiple(a, b, c int) int {
	return a + b + c
}

//BT - Return multiple int, int.
func test() (int, int) {
	return 2, 3
}

func variadic(nums ...int) {
	for i, num := range nums {
		fmt.Printf("Index: %d, value: %d\r\n", i, num)
	}
}

func main() {
	fmt.Println("Hi")

	a := add(2, 3)

	fmt.Println(a)

	b := addMultiple(2, 3, 2)

	fmt.Println(b)

	c, d := test()

	fmt.Printf("c: %d, d: %d\r\n", c, d)

	//BT - You can omit the value by using _
	_, e := test()

	fmt.Println(e)

	variadic(2, 3, 4, 5)

	fmt.Println("Passing in a slice...")

	//BT - Passing a slices with unlimited element.
	mySlices := []int{1, 2, 3, 4, 5, 6}

	//BT - This is how to pass your slice into a function.
	variadic(mySlices...)
}
