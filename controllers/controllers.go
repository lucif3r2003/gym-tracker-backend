package controllers

import (
	"gym-tracker-project/models"
	"gym-tracker-project/repositories"

	"github.com/gin-gonic/gin"
)



func SignUp(ctx *gin.Context){
	req := models.UserRequest{}
	if err := ctx.ShouldBindJSON(&req) ; err!=nil{
		ctx.JSON(500, gin.H{"error": "invalid request"})
		return
	}
	user := req.ToUser()

	//hash password
	hashedPassword, err := repositories.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "cannot hash password"})
		return
	}
	user.Password = hashedPassword

	//check dup email
	exist, err := repositories.CheckDuplicateEmail(user.Email)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "db error"})
		return
	}
	if exist{
		ctx.JSON(409, gin.H{"error": "email already exist"})
		return
	}
	
	//create user 
	errCreate := repositories.CreateUser(user)
	if errCreate != nil {
		ctx.JSON(222, gin.H{"error":"cannot create user"})
		return
	}
	ctx.JSON(200, gin.H{"message": "OK"})
}

func Login(ctx *gin.Context){
	req := models.LoginRequest{}
	if err:= ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(500, gin.H{"error":"invalid request"})
		return
	}
}





