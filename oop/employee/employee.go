package employee

import "fmt"

// //*************************************************************************************
// //BT - PART 1 - Since everything is exported. This is not a OOP
// //*************************************************************************************
// // Employee for export
// //BT - Employee - must be a letter E in order for another package can use.
// //Same as the member fields.
// type Employee struct {
// 	Firstname string
// 	Lastname  string
// }

// // PrintMe for export
// //BT - Same as the method - The first letter must be CAP
// func (e *Employee) PrintMe() {
// 	fmt.Println(e.Firstname, e.Lastname)
// }

//*************************************************************************************
//BT - PART 2 - We have to provide a public New function to initialize our private
//              employee struct.
//              Then access the public function PrintMe() in our private employee class.
//*************************************************************************************

type employee struct {
	firstname string
	lastname  string
}

// New - Act likes a constructor - Notes: It is a regular function not a method.
//BT - Act likes a constructor
func New(firstname, lastname string) employee {
	e := employee{firstname, lastname}
	return e
}

func (e *employee) PrintMe() {
	fmt.Println(e.firstname, e.lastname)
}
