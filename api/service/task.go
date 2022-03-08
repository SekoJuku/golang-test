package service

import (
	"task/api/repository"
	"task/models"
)

type TaskService struct {
	repository repository.TaskRepository
}

func NewTaskService(r repository.TaskRepository) TaskService {
	return TaskService{
		repository: r,
	}
}

func (p TaskService) Save(task models.Task) error {
	return p.repository.Save(task)
}

func (p TaskService) Update(task models.Task) error {
	return p.repository.Update(task)
}

func (p TaskService) Delete(id int64) error {
	var task models.Task
	task.ID = id
	return p.repository.Delete(task)
}

func (p TaskService) Find(task models.Task) (models.Task, error) {
	return p.repository.Find(task)
}
