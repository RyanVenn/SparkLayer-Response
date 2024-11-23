package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

var (
	toDoList   []toDo
	toDoListMx sync.Mutex
)

type toDo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {

	// Your code here
	toDoList = []toDo{}
	http.HandleFunc("/", ToDoListHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Your code here

	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodPost:
		createToDo(w, r)
	case http.MethodGet:
		getToDo(w, r)
	default:
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
	}
}

//Creates a new To-Do Task
func createToDo(w http.ResponseWriter, r *http.Request) {
	toDoListMx.Lock()
	defer toDoListMx.Unlock()

	var newToDo toDo
	err := json.NewDecoder(r.Body).Decode(&newToDo)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if newToDo.Title == "" || newToDo.Description == "" {
		http.Error(w, "Invalid input", 400)
		return
	}

	toDoList = append(toDoList, newToDo)
	w.Header().Set("Content-Type", "application/json")
	err1 := json.NewEncoder(w).Encode(newToDo)
	if err1 != nil {
		http.Error(w, err1.Error(), 500)
		return
	}
}

// Returns the current ToDoList
func getToDo(w http.ResponseWriter, r *http.Request) {
	toDoListMx.Lock()
	defer toDoListMx.Unlock()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(toDoList)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
