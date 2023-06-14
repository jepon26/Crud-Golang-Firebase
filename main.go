package main

import (
	"context"
	"fmt"
<<<<<<< HEAD
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
=======
	"text/template"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

// DBconnection establishes a connection to the database.
func DBconnection() *sql.DB {
	Driver := "mysql"
	User := ""
	Password := ""
	Name := "System"

	connection, err := sql.Open(Driver, User+":"+Password+"@tcp(127.0.0.1)/"+Name)
	if err != nil {
		panic(err.Error())
	}
	return connection
}

// templates is a global variable that holds the parsed HTML templates.
var templates = template.Must(template.ParseGlob("templates/*"))

// start is the handler function for the root URL ("/").
func start(w http.ResponseWriter, r *http.Request) {
	establishedConnection := DBconnection()

	// Inserting a record into the "employees" table.
	insertRecords, err := establishedConnection.Prepare("INSERT INTO employees(name, email) VALUES('Javier','jepon26@gmail.com')")
	if err != nil {
		panic(err.Error())
	}
	insertRecords.Exec()

	// Writing the "Hello World" message to the ResponseWriter, which sends it as the response to the client.
	templates.ExecuteTemplate(w, "start", nil)
}

// Create is the handler function for the "/create" URL.
func Create(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create", nil)
}

// main is the entry point of the application.
func main() {
	// Registering the start function to handle requests to the root URL ("/").
	http.HandleFunc("/", start)
	// Registering the Create function to handle requests to the "/create" URL.
	http.HandleFunc("/create", Create)
>>>>>>> e2ef634e75f09c51c089da1fce00dce8282711b5

	fmt.Println("Server started....")

	// Starting the HTTP server on port 8080 to listen for incoming requests.
	http.ListenAndServe(":8080", nil)
}
