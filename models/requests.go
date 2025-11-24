package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type UserRequest struct{
	Name      string             
	Password  string             
	Email     string            
	Phone     string             
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
