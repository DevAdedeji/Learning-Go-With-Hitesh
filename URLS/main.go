package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gbdbfbdhf"

func main() {
	fmt.Println("Welcome to handling URLS in Golang")
	// fmt.Println(myURL)
	// parsing url
	result, err := url.Parse(myURL)
	checkNilErr(err)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The types of query params are : %T \n", qparams)

	// fmt.Println(qparams["coursename"], qparams["paymentid"])

	// Printing all url params
	// for _, param := range qparams {
	// 	fmt.Println(param)
	// }
	// Creating a url
	partsOfURL := &url.URL{
		Scheme:  "https",
		Host:    "devadedeji.github.io",
		Path:    "/blog",
		RawPath: "user=adedeji",
	}

	anotherURL := partsOfURL.String()
	fmt.Println(anotherURL)
}

// Function to check if there is an error
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
