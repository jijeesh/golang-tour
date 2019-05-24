// This is the name of our package
package main

// These are the libraries we are going to use

import (
	// Both fmt and net are part of the Go standard libraries
	// The "net/http" library has methods to implement HTTP clients and servers
	"fmt"
	"net/http"
)

func main() {
	// The "HandleFunc" method accept a path and a function as arguments
	// (Yes,  we can pass functions as arguments, and even trat them like variable in Go)
	// However, the handler function has to have the appropriate signature (as discribed by the "handler" function below)
	http.HandleFunc("/", handler)

	// After defining our server, we finally "listen and serve" on port 8080
	// The second argumetn is the handler , which we will come to later on,
	// but for now it is left as nil,
	// and the handler defined ablov (in "HandleFunc") is used
	http.ListenAndServe(":8080", nil)
}

/*
"handler is our handler function. It has to follow the function signature of
a ResponseWriter and Request type as the arguments
"
*/

func handler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response wirter
	fmt.Fprintf(w, "Hellow World")
}
