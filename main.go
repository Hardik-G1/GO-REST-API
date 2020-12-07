package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Todo struct {
	ID     string  `json:"id"`
	Priority   string  `json:"priority"`
	Name  string  `json:"name"`
	Items *Items `json:"items"`
}

type Items struct {
	Name string `json:"name"`
	Count  string `json:"count"`
}

var TodosList []Todo


func getallTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TodosList)
}

func getsingleTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	for _, item := range TodosList {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Todo{})
}

func createNewTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newTodo Todo
	_ = json.NewDecoder(r.Body).Decode(&newTodo)
	s,_ := strconv.Atoi(TodosList[len(TodosList)-1].ID)
	newTodo.ID = strconv.Itoa(s+1)
	TodosList = append(TodosList, newTodo)
	json.NewEncoder(w).Encode(newTodo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range TodosList {
		if item.ID == params["id"] {
			TodosList = append(TodosList[:index], TodosList[index+1:]...)
			var updatedOne Todo
			_ = json.NewDecoder(r.Body).Decode(&updatedOne)
			updatedOne.ID = params["id"]
			TodosList = append(TodosList, updatedOne)
			json.NewEncoder(w).Encode(updatedOne)
			return
		}
	}
}
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range TodosList {
		if item.ID == params["id"] {
			TodosList = append(TodosList[:index], TodosList[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(TodosList)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	TodosList = append(TodosList, Todo{ID: "1", Priority: "high", Name: "add mongo db", Items: &Items{Name: "mongo", Count: "0"}})
	TodosList = append(TodosList, Todo{ID: "2", Priority: "low", Name: "host it", Items: &Items{Name: "heroku", Count: "0"}})

	// Route handles & endpoints
	r.HandleFunc("/api/TodoList", getallTodo).Methods("GET")
	r.HandleFunc("/api/TodoList/{id}", getsingleTodo).Methods("GET")
	r.HandleFunc("/api/TodoList", createNewTodo).Methods("POST")
	r.HandleFunc("/api/TodoList/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/api/TodoList/{id}", deleteTodo).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
