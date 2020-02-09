package main

import (
	"fmt"
)

// Employee - With the 'E' it means package outside can access it.
// firstname and lastname: lower case first letter - means private.
type Employee struct {
	firstname string
	lastname  string
}

func main() {

	//BT - Initialize with field name
	e := Employee{firstname: "trong", lastname: "tran"}

	fmt.Println(e.firstname)

	//BT - Init without the field name
	e1 := Employee{"trong", "tran"}

	fmt.Println(e1.lastname)

	//BT - Anonymous struct

	ns := struct {
		id  string
		age string
	}{
		id:  "1234",
		age: "24"}

	fmt.Println(ns.age)

	//BT - Slice of struct

	e2 := []Employee{
		{firstname: "trong", lastname: "tran"},
		{firstname: "jay", lastname: "tran"}}

	fmt.Println(e2[0].firstname)

	e3 := []Employee{
		{"trong", "tran"},
		{"jay", "tran"}}

	fmt.Println(e3[1].lastname)

}
