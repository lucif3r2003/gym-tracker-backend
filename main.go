package main

import (
	"fmt"
	"gym-tracker-project/database"
)

func main(){
	fmt.Println("hello project")
	database.ConnectDb()
}
