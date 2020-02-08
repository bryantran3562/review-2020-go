package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendRequest() {
	// Request (DELETE http://www.example.com/bucket/sample)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// BT - Step1: Create client
	client := &http.Client{Transport: tr}

	//BT - Step2: Create request
	data := map[string]string{"username": "admin", "password": "admin"}
	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))
	req, err := http.NewRequest("POST", "https://192.168.2.1/api/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read Response Body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}

func main() {
	fmt.Println("Hi")

	sendRequest()
}
