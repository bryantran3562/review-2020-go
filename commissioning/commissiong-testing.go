package main

import (
	"encoding/json"
	"fmt"

	//BT - Include your package
	"github.com/trong63/review-2020-go/commissioning/commissioning"
)

func main() {
	fmt.Println("Hi")
	data := map[string]string{"username": "admin", "password": "admin"}
	url := "https://192.168.2.1/api/login"
	method := "POST"
	resp, respBody, error := commissioning.DoClient(method, url, data)

	if error != nil {
		fmt.Println(error)
		return
	}

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))

	type MyJSONName struct {
		Code   int
		Result struct {
			Address    string `json:"address"`
			Permission string `json:"permission"`
			Port       string `json:"port"`
			Timestamp  string `json:"timestamp"`
			Token      string `json:"token"`
			User       string `json:"user"`
		} `json:"result"`
		Status string
	}

	var myData MyJSONName
	json.Unmarshal(respBody, &myData)

	fmt.Println(myData.Code)
	fmt.Println(myData.Result.Token)
	fmt.Println(myData.Status)

	//**************************************************************************
	//BT - GET
	//**************************************************************************

	//token := map[string]string{"token": myData.Result.Token}
	urlLoraNetwork := "https://192.168.2.1/api/loraNetwork?token=" + myData.Result.Token

	resp1, respBody1, error1 := commissioning.DoGet(urlLoraNetwork)

	if error1 != nil {
		fmt.Println(error)
		return
	}

	// Display Results
	fmt.Println("response Status : ", resp1.Status)
	fmt.Println("response Headers : ", resp1.Header)
	fmt.Println("response Body : ", string(respBody1))

	//BT - Because the 'respBody1' is a []byte. So, we need to convert it to json for easy access the element.
	//     However, the key is a string, but the value is a different type - i.e - int, string..etc. This is why
	//     we have to setup a variable with: map[string]interface{}
	var myJONdata map[string]interface{}
	//BT - Then call Unmarshall to convert from []bytes to our map[string]interface{}
	json.Unmarshal(respBody1, &myJONdata)
	//BT - So now we can access it with [key]interface{}
	for k, v := range myJONdata {
		fmt.Println("k:", k, "v:", v)
	}

	//BT - The problem is that - when a key has a value as a map. Then how do you access the data in map.
	//myJONdata["result"]: is a key
	//.(map[string]interface)
	// We then have to set it up with form below.
	b := myJONdata["result"].(map[string]interface{})

	//BT - Now we can access the map[string]interface{}
	fmt.Println(b["__v"])

	//BT - Setup another map nested in the key above
	c := b["addressRange"].(map[string]interface{})

	fmt.Println(c["end"])

}
