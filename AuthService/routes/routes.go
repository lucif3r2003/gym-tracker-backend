package routes

import (
	"gym-tracker-project/controllers"

	"github.com/gin-gonic/gin"
)

func IdentityRoute(identityRoute *gin.Engine){
	identityRoute.POST("/users/signup", controllers.SignUp)
	identityRoute.POST("/users/login", controllers.Login)
}
