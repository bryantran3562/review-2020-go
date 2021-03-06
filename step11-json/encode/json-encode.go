package main

import (
	"encoding/json"
	"fmt"
)

type myData struct {
	Name string
	Age  int
}

func main() {

	data := myData{Name: "trong", Age: 24}

	//BT - Take a Go struct and encoded into the json format. This will be a decimal bytes.
	//     For example, fmt.Println(dataString)
	dataString, _ := json.Marshal(data)

	fmt.Printf("%s\n", dataString)

	//BT - Passing in two slices

	dataString2, _ := json.Marshal([]myData{data, data})

	fmt.Println(string(dataString2))

	// //BT - map

	// test := map[string]interface{}{

	// 	"birdSounds": map[string]string{
	// 		"pigeon": "coo",
	// 		"eagle":  "squak",
	// 	},
	// 	"total birds": 2,
	// }

	// dataString3, _ := json.Marshal(test)

	// fmt.Println(string(dataString3))

}
