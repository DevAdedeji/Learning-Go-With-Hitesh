package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("API")
	r := mux.NewRouter()
	// seedling
	courses = append(courses, Course{CourseId: "1", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "Go Crash Course", CoursePrice: 99, Author: &Author{Fullname: "Hitesh Choudary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "VueJS Crash Course", CoursePrice: 399, Author: &Author{Fullname: "Hitesh Choudary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MEVN", CoursePrice: 999, Author: &Author{Fullname: "Hitesh Choudary", Website: "lco.dev"}})

	// Listen to port
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))

}

// controllers -file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by DevAdedeji</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// Loop through courses, find matching id and reutn respone
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found wth given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No Data")
		return
	}

	// Check only if title is duplicate
	for _, loopcourse := range courses {
		if course.CourseName == loopcourse.CourseName {
			json.NewEncoder(w).Encode("Course already exist")
			break
		} else {
			// generate a unique ID, string
			rand.Seed(time.Now().UnixNano())
			course.CourseId = strconv.Itoa(rand.Intn(100))
			// append course into courses
			courses = append(courses, course)
			json.NewEncoder(w).Encode(courses)
		}
	}

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")
	//  grab id from request
	params := mux.Vars(r)
	// loop id, remove, add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// send a response when id is not found
	json.NewEncoder(w).Encode("No course found with that given ID")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course is deleted")
			break
		} else {
			json.NewEncoder(w).Encode("Course doesnt exist")
		}
	}
}
