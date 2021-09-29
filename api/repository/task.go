package repository

import (
	"task/infrastructure"
	"task/models"
)

//TaskRepository -> TaskRepository
type TaskRepository struct {
	db infrastructure.Database
}

// NewTaskRepository : fetching database
func NewTaskRepository(db infrastructure.Database) TaskRepository {
	return TaskRepository{
		db: db,
	}
}

//Save -> Method for saving task to database
func (t TaskRepository) Save(task models.Task) error {
	return t.db.DB.Create(&task).Error
}

//Update -> Method for updating Task
func (t TaskRepository) Update(task models.Task) error {
	return t.db.DB.Save(&task).Error
}

//Find -> Method for fetching task by id
func (t TaskRepository) Find(task models.Task) (models.Task, error) {
	var tasks models.Task
	err := t.db.DB.
		Debug().
		Model(&models.Task{}).
		Where(&task).
		Take(&tasks).Error
	return tasks, err
}

//Delete Deletes Task
func (t TaskRepository) Delete(task models.Task) error {
	return t.db.DB.Delete(&task).Error
}
