package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"text/template"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// templates is a global variable that holds the parsed HTML templates.
var templates = template.Must(template.ParseGlob("templates/*"))

// start is the handler function for the root URL ("/").
func start(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	config := &firebase.Config{
		DatabaseURL: "https://your-firebase-project.firebaseio.com",
	}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}

	// Firebase Realtime Database reference
	dbClient, err := app.Database(ctx)
	if err != nil {
		log.Fatalf("Failed to get database client: %v", err)
	}

	// Example: Writing data to Firebase Realtime Database
	ref := dbClient.NewRef("employees")
	err = ref.Set(ctx, map[string]interface{}{
		"name":  "Javier",
		"email": "jepon26@gmail.com",
	})
	if err != nil {
		log.Fatalf("Failed to write data: %v", err)
	}

	// Writing the "Hello World" message to the ResponseWriter, which sends it as the response to the client.
	templates.ExecuteTemplate(w, "start", nil)
}

// create is the handler function for the "/create" URL.
func create(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create", nil)
}

// main is the entry point of the application.
func main() {
	// Register the handler functions
	http.HandleFunc("/", start)
	http.HandleFunc("/create", create)

	// Start the HTTP server
	fmt.Println("Server listening on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}