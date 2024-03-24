package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "Welcome to the website")
	})

	// this is where we would serve static images/css/etc
	fileServer := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	fmt.Println("Now listening on port 80")
	http.ListenAndServe(":80", nil)
}