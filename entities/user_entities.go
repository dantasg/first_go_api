package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Email    string
	Password string
}

type APIUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type UserLogin struct {
	Email    string
	Password string
}
