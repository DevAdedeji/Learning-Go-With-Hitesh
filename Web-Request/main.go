package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web Requests")
	// Making a get request to the url
	response, err := http.Get(url)
	checkNilErr(err)
	// Type of Response
	fmt.Printf("Response is of type: %T", response)
	// Its important to close all calls to API
	defer response.Body.Close()
	databyte, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	// Print response in a string format
	content := string(databyte)
	fmt.Println(string(content))
}

// Function to check if there is an error
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
