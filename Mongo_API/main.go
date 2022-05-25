package main

import (
	"Go-with-Hitesh/Mongo_API/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")
	// From the routers folder
	r := routers.Router()
	fmt.Println("Server is getting started...")
	// Server
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Listening at port 8080....")
}
