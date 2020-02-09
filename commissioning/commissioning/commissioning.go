package commissioning

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DoClient Client Connection
func DoClient(method string, url string, data map[string]string) (*http.Response, []byte, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// BT - Step1: Create client
	client := &http.Client{Transport: tr}

	//BT - Step2: Create request

	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-type", "application/json")
	fmt.Println(req.URL)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	defer resp.Body.Close()

	// Read Response Body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	return resp, respBody, nil
}

// DoGet get connection
func DoGet(url string) (*http.Response, []byte, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// BT - Step1: Create client
	client := &http.Client{Transport: tr}

	//BT - Step2: Create request
	req, err := client.Get(url)
	req.Header.Set("Content-type", "application/json")
	fmt.Println(req.Request.URL)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	defer req.Body.Close()

	// Read Response Body
	respBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	return req, respBody, nil
}
