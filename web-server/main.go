package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Apa Kabar!!")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ini Halaman About")
}

func withHtml(w http.ResponseWriter, r *http.Request) {
	var data = map[string]string{
		"Name":    "Rijal Arfani",
		"Message": "Belajar Golang",
	}

	var t, err = template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Word")
	})

	http.HandleFunc("/index", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/person", withHtml)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
