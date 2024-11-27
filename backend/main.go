package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type toDoItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// In memory variable to hold list items
var toDoList []toDoItem

func main() {
	http.HandleFunc("/", ToDoListHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting function")
	}
}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	// Allow cross origin requests
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Your code here
	if r.Method == http.MethodOptions {
		// Allows CORS preflight request
		return

	} else if r.Method == http.MethodPost {
		// Allow user to create new item
		var newItem toDoItem
		// Decode item details from request
		err := json.NewDecoder(r.Body).Decode(&newItem)
		// Check all data is valid and return an error if not
		if err != nil || newItem.Title == "" || newItem.Description == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// If request valid add item to list
		toDoList = append(toDoList, newItem)
		w.WriteHeader(http.StatusCreated)

	} else if r.Method == http.MethodGet {
		// Allow user to get all items in list
		// Turn list into JSON and add to response
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(toDoList)
		// If error occurs change response status code
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	} else {
		// Don't allow any other methods
		fmt.Fprintf(w, "Invalid method")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
