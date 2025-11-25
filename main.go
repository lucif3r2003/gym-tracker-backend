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
	fmt.Println("hello project")
	database.ConnectDb()
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	
	//set up gin -_- blyat
	r := gin.Default()
	routes.IdentityRoute(r)	
	r.Run(port)
}
