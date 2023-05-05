package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	Nim     string
	Nama    string
	Address string
}

var students = []student{
	{"1804004", "Rijal Arfani", "Tembok Banjaran"},
	{"1804001", "Panggih", "Janinegara"},
	{"1804002", "Balqis Shafa Wardahni", "Gumalar"},
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(students)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func getUserByNim(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var nim = r.FormValue("nim")
		var result []byte
		var err error

		for _, student := range students {
			if student.Nim == nim {
				result, err = json.Marshal(student)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func main() {

	http.HandleFunc("/students", getAllUsers)
	http.HandleFunc("/student", getUserByNim)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
