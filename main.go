package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	
	"github.com/gorilla/mux"
)

// Todo -
type Todo struct {
	Name string
	Completed bool
	Due time.Time
}

// Todos -
type Todos []Todo



func main() {

	// create a router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	// handle fatal errors and serve the routes on 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index - 
func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Welcome!")
}

// TodoIndex -
func TodoIndex(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Todo Index!")
}

// TodoShow -
func TodoShow(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)  // get the vars
	todoID := vars["todoID"]
	fmt.Fprintln(w, "Todo show:", todoID)
}