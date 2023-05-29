package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)


var templates = template.Must(template.ParseGlob("templates/*"))

func main(){
	http.HandleFunc("/", start)
    log.Println("Server started....")
	http.ListenAndServe(":8080", nil)
}


func start(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}

