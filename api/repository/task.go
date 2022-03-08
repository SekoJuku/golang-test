package repository

import (
	"task/infrastructure"
	"task/models"
)

type TaskRepository struct {
	db infrastructure.Database
}

func NewTaskRepository(db infrastructure.Database) TaskRepository {
	return TaskRepository{
		db: db,
	}
}

func (t TaskRepository) Save(task models.Task) error {
	return t.db.DB.Create(&task).Error
}

func (t TaskRepository) Update(task models.Task) error {
	return t.db.DB.Save(&task).Error
}

func (t TaskRepository) Find(task models.Task) (models.Task, error) {
	var tasks models.Task
	err := t.db.DB.
		Debug().
		Model(&models.Task{}).
		Where(&task).
		Take(&tasks).Error
	return tasks, err
}

func (t TaskRepository) Delete(task models.Task) error {
	return t.db.DB.Delete(&task).Error
}
