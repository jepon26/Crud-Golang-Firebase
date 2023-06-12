package main

import (
	"database/sql"
	"fmt"
	"text/template"

	//"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)


func DBconnection()(connection *sql.DB){
	Driver:= "mysql"
	User:= "add User"
	Password:="Add Password"
	Name:="System"


	connection, err:= sql.Open(Driver, User+":"+Password+"@tcp(127.0.0.1)/"+Name)
	if err != nil {
		panic(err.Error())
	}
	return connection
}




// templates is a global variable that holds the parsed HTML templates.
var templates = template.Must(template.ParseGlob("templates/*"))

// main is the entry point of the application.
func main() {
	// Registering the start function to handle requests to the root URL ("/").
	http.HandleFunc("/", start)
	http.HandleFunc("/create", Create)

	fmt.Println("Server started....")
	
	// Starting the HTTP server on port 8080 to listen for incoming requests.
	http.ListenAndServe(":8080", nil)
}

// start is the handler function for the root URL ("/").
func start(w http.ResponseWriter, r *http.Request) {

	establishedConnection:= DBconnection()

	insertRecords,err:= establishedConnection.Prepare("INSERT INTO employees(name, email) VALUES('Javier','jepon26@gmail.com')")

      if err!=nil{
		panic(err.Error())
	  }
	  insertRecords.Exec()
	  

	// Writing the "Hello World" message to the ResponseWriter, which sends it as the response to the client.
	templates.ExecuteTemplate(w, "start", nil)
}

func Create(w http.ResponseWriter, r*http.Request){
	templates.ExecuteTemplate(w, "create", nil)
}
