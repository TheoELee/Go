package main

import (
	"fmt"
	"net/http"
)


func main() {
	http.HandleFunc("/", func (response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "Hello, you've requested: %s\n", request.URL.Path)
	})

	fmt.Println("now listening on port 80...")
	http.ListenAndServe(":80", nil)
}
