package service

import (
	"task/api/repository"
	"task/models"
)

//TaskService TaskService struct
type TaskService struct {
	repository repository.TaskRepository
}

//NewTaskService : returns the TaskService struct instance
func NewTaskService(r repository.TaskRepository) TaskService {
	return TaskService{
		repository: r,
	}
}

//Save -> calls task repository save method
func (p TaskService) Save(task models.Task) error {
	return p.repository.Save(task)
}

// Update -> calls taskrepo update method
func (p TaskService) Update(task models.Task) error {
	return p.repository.Update(task)
}

// Delete -> calls task repo delete method
func (p TaskService) Delete(id int64) error {
	var task models.Task
	task.ID = id
	return p.repository.Delete(task)
}

// Find -> calls task repo find method
func (p TaskService) Find(task models.Task) (models.Task, error) {
	return p.repository.Find(task)
}
