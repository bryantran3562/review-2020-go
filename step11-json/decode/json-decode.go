package main

import (
	"fmt"
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

	// data := `{
	// 	"name" : "trong",
	// 	"age" : 42
	// }`

	// //BT - The rule is that - var your_variable_name type. The interface{} will accept any type.
	// var myMap map[string]interface{}

	// //BT - The json.Unmarshal convert string to Go map

	// //BT - Because it is a map ==> access it using [key]
	// error := json.Unmarshal([]byte(data), &myMap)

	// if error != nil {
	// 	log.Fatalln("Could not convert to map...")
	// }

	// fmt.Println(myMap["age"])

	//BT - Case 1

	// data1 := `{
	// 	"species": "horse",
	// 	"description" : "race likes a horse"
	// }`

	// type myData1Struct struct {
	// 	Species     string
	// 	Description string
	// }

	// var myData1 myData1Struct

	// //BT - Because loading the string into our struct.
	// json.Unmarshal([]byte(data1), &myData1)

	// //BT - We then access our struct member using the dot.
	// fmt.Println(myData1.Species)

	// //BT - Case 2 - Array
	// //Look at the key, value and create a struct to have the same key, value.
	// //If the key, value are the same, create an array of your struct.

	// data2 := `[
	// 				{
	// 					"species": "pigeon",
	// 					"description": "likes to perch on rocks"
	// 				},
	// 				{
	// 					"species":"eagle",
	// 					"description":"bird of prey"
	// 				}
	// ]`

	// var myData2 []myData1Struct

	// //BT - Load our array of string into our array of struct.
	// json.Unmarshal([]byte(data2), &myData2)
	// //BT - Accessing array of struct with dot for member
	// fmt.Println(myData2[0].Species)

	// //BT - Case 3: Embedded objects
	// //Just create another struct to represent the embedded object and then nest it into your outer struct.

	// data3 := `{
	// 	"species": "pigeon",
	// 	"decription": "likes to perch on rocks",
	// 	"dimensions": {
	// 	  "height": 24,
	// 	  "width": 10
	// 	}
	// }`

	// type Dimensions struct {
	// 	Height int
	// 	Width  int
	// }

	// type myData3Struct struct {
	// 	Species     string
	// 	Description string
	// 	Dimensions  Dimensions
	// }

	// var myData3 myData3Struct

	// //BT - Loading our string into our struct with another struct nested. Access it using dot.
	// json.Unmarshal([]byte(data3), &myData3)

	// fmt.Println(myData3.Dimensions.Height)

	// //BT - Case 4 - Unstructure

	// data4 := `{
	// 	"birds": {
	// 	  "pigeon":"likes to perch on rocks",
	// 	  "eagle":"bird of prey"
	// 	},
	// 	"animals": "none"
	//   }`

	// var myMap1 map[string]interface{}

	// json.Unmarshal([]byte(data4), &myMap1)

	// fmt.Println(myMap1["birds"])

	// //BT - Accessing nested map[string]interface{}

	// b := myMap1["birds"].(map[string]interface{})

	// fmt.Println(b["eagle"])
}
