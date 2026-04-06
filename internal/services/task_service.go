package services

import "github.com/VJunqui/go-todo-api/internal/models"

type TaskService struct {
	tasks  []models.Task
	nextID int
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks:  []models.Task{},
		nextID: 1,
	}
}

func (s *TaskService) GetAll() []models.Task {
	return s.tasks
}

func (s *TaskService) Create(task models.Task) models.Task {
	task.ID = s.nextID
	s.nextID++
	task.Done = false

	s.tasks = append(s.tasks, task)
	return task
}

func (s *TaskService) Update(updated models.Task) (models.Task, bool) {
	for i, t := range s.tasks {
		if t.ID == updated.ID {
			s.tasks[i].Done = updated.Done
			return s.tasks[i], true
		}
	}
	return models.Task{}, false
}

func (s *TaskService) Delete(id int) bool {
	for i, t := range s.tasks {
		if t.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return true
		}
	}
	return false
}