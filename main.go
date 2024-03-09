package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) { // r *http.Request is a pointer to the request object
	if err := r.ParseForm(); err != nil { // ParseForm() parses the raw query from the URL and updates r.Form
		fmt.Fprintf(w, "ParseForm() err: %v", err) // w is the response writer
		return // return to exit the function
	}
	fmt.Fprintf(w, "POST request successful\n") // write to the response writer
	name := r.FormValue("name") // r.FormValue() returns the first value for the named component of the query
	address := r.FormValue("address") // r.FormValue() returns the first value for the named component of the query
	fmt.Fprintf(w, "Name = %s\n", name) // write to the response writer
	fmt.Fprintf(w, "Address = %s\n", address) 

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!\n")

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
