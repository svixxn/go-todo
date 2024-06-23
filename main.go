package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Task struct {
	Text string	`json:"text"`
	Priority int	`json:"priority"`
	IsCompleted bool 	`json:"isCompleted"`
	AuthorId int	`json:"authorId"`
}

var tasks []Task


func handleTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(tasks)
	}

	if r.Method == "POST" {
		var newTask Task
		err := json.NewDecoder(r.Body).Decode(&newTask)
    	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    	}
		tasks = append(tasks, newTask)
		json.NewEncoder(w).Encode(newTask)
	}
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/task", handleTask)
	log.Fatal(http.ListenAndServe(":8080", mux))
}