package main

import (
	"log"
	"net/http"
	"github.com/VJunqui/go-todo-api/internal/handlers"
	"github.com/VJunqui/go-todo-api/internal/services"
)

func main() {
	taskService := services.NewTaskService()
	taskHandler := handlers.NewTaskHandler(taskService)

	http.HandleFunc("/tasks", taskHandler.HandleTasks)

	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}