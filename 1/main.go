package main

import (
	"fmt"
	"net/http"
)

func formHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprint(w, "<h1>form!!!!!!!!!!</h1>")
	} else {
		fmt.Printf("wrong method")
	}
	err := r.ParseForm()

	if err != nil {
		fmt.Print("unable to parse the form")
		return
	}
	name := r.FormValue("name-data")
	fmt.Fprint(w, "name is %s ", name)

}

func helloHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		fmt.Fprint(w, "<h1>hello!!!!!!!!</h1>")
	} else {
		fmt.Printf("wrong method")
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("/home/arjun/golang/temp"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandle)
	http.HandleFunc("/hello", helloHandle)
	http.ListenAndServe(":9012", nil)
}
