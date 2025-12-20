package main

import (
	"fmt"
	"gym-tracker-project/database"
	"gym-tracker-project/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	//get port and connect db

	godotenv.Load(".env")
	fmt.Println("hello project")
	database.ConnectDb()
	port := os.Getenv("PORT")
	
	//set up gin -_- blyat
	r := gin.Default()
	routes.IdentityRoute(r)	
	r.Run(port)
}
