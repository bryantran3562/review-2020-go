package main

import (
	"fmt"
)

type myStruct struct {
	name string
	age  int
}

//BT - Method associate with the above struct with pointer
func (m *myStruct) printMePointer(aNewName string) {
	m.name = aNewName
}

//BT - Method associate with the above struct with copy by value
func (m myStruct) printMeCopy(aNewName string) {
	m.name = aNewName
}

func main() {

	fmt.Println("Hi")

	s := myStruct{name: "trong", age: 42}

	fmt.Println(s.name, s.age)

	s.printMePointer("hary")

	//Important
	//BT - After printMePointer - With the pointer receiver, the change will reflect to the caller.

	fmt.Println(s.name)

	s.printMeCopy("micheal")

	//Important
	//BT - After printMeCopy - With value receiver, the change only effect the local but not to the caller.

	fmt.Println(s.name)
}
