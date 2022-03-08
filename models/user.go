package models

import "time"

type User struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"username" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u *User) TableName() string {
	return "users"
}

type UserLogin struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UserRegister struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
	Name     string `form:"name"`
	Surname  string `form:"surname"`
}

func (u *User) ResponseMap(user *User) map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = user.ID
	resp["name"] = user.Name
	resp["surname"] = user.Surname
	resp["email"] = user.Email
	resp["password"] = user.Password
	return resp
}
