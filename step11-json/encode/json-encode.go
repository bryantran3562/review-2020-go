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

	dataString, _ := json.Marshal(data)

	fmt.Println(string(dataString))

	//BT - Passing in two slices

	dataString2, _ := json.Marshal([]myData{data, data})

	fmt.Println(string(dataString2))

	//BT - map

	test := map[string]interface{}{

		"birdSounds": map[string]string{
			"pigeon": "coo",
			"eagle":  "squak",
		},
		"total birds": 2,
	}

	dataString3, _ := json.Marshal(test)

	fmt.Println(string(dataString3))

}
