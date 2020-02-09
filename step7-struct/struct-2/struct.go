package main

import "fmt"

//BT - Nested struct
type employee struct {
	firstname string
	lastname  string
	address   address
}

type address struct {
	address string
	phone   string
}

func main() {

	fmt.Println("Hi")

	//BT - Initialize the struct and nested struct
	ns := employee{
		firstname: "trong",
		lastname:  "tran",
		address: address{
			address: "123 street CA 44323",
			phone:   "123-456-7890"}}

	fmt.Println(ns.address.address)

}
