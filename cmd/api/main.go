package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/VJunqui/go-todo-api/internal/models"
)

var tasks []models.Task
var nextID = 1

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)

	case http.MethodPost:
		var task models.Task

		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		task.ID = nextID
		nextID++
		task.Done = false

		tasks = append(tasks, task)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)

	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/tasks", tasksHandler)

	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
