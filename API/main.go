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

//!IMPORTANT POSTMAN COLLECTION LINK
//*INFORMATION: https://www.getpostman.com/collections/ab8c9976f673a04c4612

// MODEL FOR COURSE - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// FAKE DATABASE
var courses []Course

// middleware, helper - file

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	r := mux.NewRouter()

	courses = append(courses, Course{
		CourseId: "2",
		CourseName: "ReactJS",
		CoursePrice: 299,
		Author: &Author{
			FullName: "Gaurav",
			Website: "gaurav@google.com",
		},
	})

	courses = append(courses, Course{
		CourseId: "4",
		CourseName: "MERN Stack",
		CoursePrice: 199,
		Author: &Author{
			FullName: "Gaurav",
			Website: "gaurav@go.dev",
		},
	})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/getAllCourses", getAllCourses).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", getCourseById).Methods("GET")
	r.HandleFunc("/course/{id}", updateCourseById).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourseById).Methods("DELETE")

	// listen to the port

	log.Fatal(http.ListenAndServe(":4000", r))

	http.Handle("/", r)
}

// CONTROLLERS - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET all courses")

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(courses)

}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET One course")

	w.Header().Set("Content-type", "application/json")

	// GRAB ID FORM REQUEST

	params := mux.Vars(r)

	fmt.Println("Param", params)

	// Loop through the courses and find the matching id and return the response
	for _, course := range courses {

		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")

	w.Header().Set("Content-type", "application/json")

	// what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Body is empty")
	}

	// what if empty object - {}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("EMPTY OBJECT")
		return
	}

	// generate uniqueID, into string
	// append new course in to courses

	rand.Seed(time.Now().UnixNano())

	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}

func updateCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")

	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	fmt.Println("Param ", params)
	
	// Loop through the courses and find the matching id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			fmt.Println("Sorted score : ", courses)
			
			var newCourse Course
			
			_ = json.NewDecoder(r.Body).Decode(&newCourse)
			newCourse.CourseId = params["id"]
			courses = append(courses, newCourse)
			json.NewEncoder(w).Encode(newCourse)
			break //OR return
		}
	}	
}

func deleteCourseById(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete one course")

	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	fmt.Println("Param", params)
	
	// Loop through the courses and find the matching id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)			
			json.NewEncoder(w).Encode("Course deleted successfully")
			break //OR return
		}
	}
}
