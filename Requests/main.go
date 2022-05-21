package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const myURL = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fmt.Println("Welcome to Get request")
	performPostJSONRequest(myURL)
}

func performPostJSONRequest(myURL string) {
	// Fake JSON payload
	requestBody := strings.NewReader(`
	{
		"userId": 1,
		"title": "GoLang",
		"body": "Performing Post Request with Golang"
	}
	`)

	result, err := http.Post(myURL, "application/json", requestBody)
	// checking if there is any error
	checkNilErr(err)
	//  close all calls
	defer result.Body.Close()
	//using ioutil package to read the response
	data, err := ioutil.ReadAll(result.Body)
	// checking if there is any error
	checkNilErr(err)
	// print response as a string
	fmt.Println(string(data))
}

// If there is any err, exit!!!
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
