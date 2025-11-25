package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type UserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Phone    string `json:"phone"`
}

func (r *UserRequest) ToUser() User {
	return User{
		ID:        primitive.NewObjectID(),
		Name:      r.Name,
		Password:  r.Password, 
		Email:     r.Email,
		Phone:     r.Phone,
		Create_At: time.Now(),
		Update_At: time.Now(),
	}
}
