package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"text/template"
)

// templates is a global variable that holds the parsed HTML templates.
var templates = template.Must(template.ParseGlob("templates/*"))

// main is the entry point of the application.
func main() {
	ctx := context.Background()

	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	config := &firebase.Config{
		DatabaseURL: "https://your-firebase-project.firebaseio.com",
	}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}

	// Registering the start function to handle requests to the root URL ("/").
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "create", nil)
	})

	fmt.Println("Server started....")

	// Starting the HTTP server on port 8080 to listen for incoming requests.
	http.ListenAndServe(":8080", nil)
}
