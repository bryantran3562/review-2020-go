package main

import (
	"fmt"
)

type myStruct struct {
	name string
	age  int
}

//BT - Method associate with the above struct with pointer
func (m *myStruct) printMePointer() string {
	return m.name
}

//BT - Method associate with the above struct with copy
func (m myStruct) printMeCopy() string {
	return m.name
}

func main() {

	fmt.Println("Hi")

	s := myStruct{name: "trong", age: 42}

	fmt.Println(s.name, s.age)

	fmt.Println(s.printMePointer())

	fmt.Println(s.printMeCopy())
}
