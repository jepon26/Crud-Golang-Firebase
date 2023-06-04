package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// templates is a global variable that holds the parsed HTML templates.
var templates = template.Must(template.ParseGlob("templates/*"))

// main is the entry point of the application.
func main() {
	// Registering the start function to handle requests to the root URL ("/").
	http.HandleFunc("/", start)
	log.Println("Server started....")
	
	// Starting the HTTP server on port 8080 to listen for incoming requests.
	http.ListenAndServe(":8080", nil)
}

// start is the handler function for the root URL ("/").
func start(w http.ResponseWriter, r *http.Request) {
	// Writing the "Hello World" message to the ResponseWriter, which sends it as the response to the client.
	fmt.Fprintf(w, "Hello World")
}

