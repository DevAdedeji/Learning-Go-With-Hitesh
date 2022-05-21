package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const myURL = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	title := r.FormValue("title")
	body := r.FormValue("body")
	data := url.Values{}
	data.Add("title", title)
	data.Add("body", body)

	result, err := http.PostForm(myURL, data)
	if err != nil {
		panic(err)
	}
	defer result.Body.Close()
	content, err := ioutil.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
