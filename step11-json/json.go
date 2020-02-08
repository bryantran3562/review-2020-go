package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hi")

	// data := `{
	// 	"name" : "trong",
	// 	"age" : "42"
	// }`

	// //BT - The rule is that - var your_variable_name type
	// var myMap map[string]string

	// //BT - The json.Unmarshal convert string to Go map
	// error := json.Unmarshal([]byte(data), &myMap)

	// if error != nil {
	// 	log.Fatalln("Could not convert to map...")
	// }

	// fmt.Println(myMap)

	// for k, v := range myMap {
	// 	fmt.Printf("Key: %s, Value: %s\r\n", k, v)
	// }

	//BT - The issue above is that - you can not set age to a number like 42.

	data := `{
		"name" : "trong",
		"age" : 42
	}`

	//BT - The rule is that - var your_variable_name type. The interface{} will accept any type.
	var myMap map[string]interface{}

	//BT - The json.Unmarshal convert string to Go map
	error := json.Unmarshal([]byte(data), &myMap)

	if error != nil {
		log.Fatalln("Could not convert to map...")
	}

	fmt.Println(myMap["age"])

}
