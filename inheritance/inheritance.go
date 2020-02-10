package main

import (
	"fmt"
)

// Employee - Employee
type Employee struct {
	name string
}

// PrintEmployeeName - Print employee name
func (e *Employee) PrintEmployeeName() {
	fmt.Println("Employee name:", e.name)
}

//*******************************************************************************
//BT - Inheritance - Using composite or embedded struct for inheritance
//*******************************************************************************

// EngDept - Eng Dept
type EngDept struct {
	salary   int
	employee Employee
}

// PrintEngDept - Print the Eng
func (ed *EngDept) PrintEngDept() {
	fmt.Printf("Employee name: %s - salary: %d", ed.employee.name, ed.salary)
	fmt.Println()
}

func main() {

	fmt.Println("hi")

	ed := EngDept{salary: 65000,
		employee: Employee{"trong tran"}}

	ed.PrintEngDept()

	//BT - Access Parent method

	ed.employee.PrintEmployeeName()
}
