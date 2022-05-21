package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON form")
	DecodeJSON()
}

func EncodeJSON() {
	myCourse := []course{
		{"ReactJS Bootcamp", 299, "Udemy", "adedejimayowa9233", []string{"Web-dev", "js"}},
		{"MERN Bootcamp", 199, "Udemy", "mayowa9233", []string{"Web-dev", "js"}},
		{"Angular Bootcamp", 299, "Udemy", "adedeji9233", nil},
	}

	// package data as JSON

	finalJSON, err := json.MarshalIndent(myCourse, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJSON)

}

func DecodeJSON() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "ReactJS Bootcamp",
			"Price": 299,
			"website": "Udemy",
			"tags": ["Web-dev","js"]
		}
	`)

	var lcoCourse course
	checkvalid := json.Valid(jsonDataFromWeb)
	if checkvalid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON wasnt valid")
	}

	// some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}
}
