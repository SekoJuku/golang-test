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

//Save -> calls post repository save method
func (p TaskService) Save(post models.Task) error {
	return p.repository.Save(post)
}

// Update -> calls postrepo update method
func (p TaskService) Update(post models.Task) error {
	return p.repository.Update(post)
}

// Delete -> calls post repo delete method
func (p TaskService) Delete(id int64) error {
	var post models.Task
	post.ID = id
	return p.repository.Delete(post)
}

// Find -> calls post repo find method
func (p TaskService) Find(post models.Task) (models.Task, error) {
	return p.repository.Find(post)
}
