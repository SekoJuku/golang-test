package models

import "time"

type Task struct {
	ID          int64     `gorm:"primary_key;auto_increment" json:"id"`
	Name        string    `gorm:"size:200" json:"name"`
	Description string    `gorm:"size:3000" json:"description" `
	Author      string    `gorm:"size:200" json:"author"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// TableName method sets table name for Post model
func (post *Task) TableName() string {
	return "tasks"
}

//ResponseMap -> response map method of Post
func (post *Task) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = post.ID
	resp["name"] = post.Name
	resp["description"] = post.Description
	resp["author"] = post.Author
	return resp
}
