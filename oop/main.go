package main

import (
	"github.com/trong63/review-2020-go/oop/employee"
)

func main() {

	//*************************************************************************************
	//BT - PART 1 - Since everything is exported. This is not a OOP
	//*************************************************************************************

	// e := employee.Employee{Firstname: "trong", Lastname: "tran"}

	// e.PrintMe()

	//*************************************************************************************
	//BT - PART 2 - We have to provide a public New function to initialize our private
	//              employee struct.
	//              Then access the public function PrintMe() in our private employee class.
	//*************************************************************************************
	e := employee.New("trong", "tran")

	e.PrintMe()

}
