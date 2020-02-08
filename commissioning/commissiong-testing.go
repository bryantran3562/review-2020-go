package main

import (
	"fmt"
	//BT - Include your package
	"github.com/trong63/review-2020-go/commissioning/commissioning"
)

func main() {
	fmt.Println("Hi")
	//BT - Package Name.Struct Name
	c := commissioning.Commissioning{"admin", "admin", "192.168.2.1"}
}
