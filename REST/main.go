package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type courseInfo struct {
	TITLE string `json:"Title"`
}

var courses map[string]courseInfo

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the REST API!!!")
}

func allcourses(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of all courses\n")
	kv := r.URL.Query()

	for k, v := range kv {
		fmt.Println(k, v)
	}
	json.NewEncoder(w).Encode(courses)
}

func course(w http.ResponseWriter, r *http.Request) {
	/*
		params := mux.Vars(r)
		fmt.Fprintf(w, "Detail of course "+params["courseid"])
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, r.Method)
	*/
	params := mux.Vars(r)

	if r.Header.Get("Content-Type") == "application/json" {

		//POST to create a new course
		if r.Method == "POST" {
			reqBody, err := io.ReadAll(r.Body)

			if err == nil {
				// Convert JSON to object using Unmarshal
				var newCourse courseInfo
				json.Unmarshal(reqBody, &newCourse)

				if newCourse.TITLE == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte("422 - Please supply course information in JSON format"))
					return
				}

				// check if course exists; add only if course does not exist
				params := mux.Vars(r)
				if _, ok := courses[params["courseid"]]; !ok {
					courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Course added: " + params["courseid"]))
				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte("409 - Duplicate course ID"))
				}
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply course information in JSON format"))
			}
		}
	}

	//Update
	if r.Method == "PUT" {
		var newCourse courseInfo
		reqBody, err := io.ReadAll(r.Body)

		if err == nil {
			json.Unmarshal(reqBody, &newCourse)

			if newCourse.TITLE == "" {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply course information in JSON format"))
				return
			}

			params := mux.Vars(r)
			if _, ok := courses[params["courseid"]]; !ok {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - No such course"))
			} else {
				courses[params["courseid"]] = newCourse
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte("201 - Course added: " + params["courseid"]))
			}
		} else {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("422 - Please supply course information in JSON format"))
		}
	}

	//Retreive
	if r.Method == "GET" {
		if _, ok := courses[params["courseid"]]; ok {
			json.NewEncoder(w).Encode(courses[params["courseid"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No such course"))
		}
	}

	//Delete
	if r.Method == "DELETE" {
		if _, ok := courses[params["courseid"]]; ok {
			delete(courses, params["courseid"])
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No such course"))
		}
	}

}

func main() {
	courses = make(map[string]courseInfo)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1", home)

	router.HandleFunc("/api/v1/courses", allcourses)
	router.HandleFunc("/api/v1/courses/{courseid}", course).Methods(
		"GET", "POST", "PUT", "DELETE")

	fmt.Println("Listening at port 5050")
	log.Fatal(http.ListenAndServe(":5050", router))
}
