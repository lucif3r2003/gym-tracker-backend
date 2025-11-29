package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DB *mongo.Database

func ConnectUrl() string {
	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	cluster := os.Getenv("MONGO_CLUSTER")
	db := os.Getenv("MONGO_DB")
	if user == "" || password == "" || cluster == "" || db == "" {
		log.Fatal("Mongo environment variables not fully set")
	}

	// Build URI
	return fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", user, password, cluster, db)
}

func ConnectDb(){
	connectUrl := ConnectUrl()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectUrl).SetServerAPIOptions(serverAPI)

	client, err:= mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")


	godotenv.Load(".env")
	DB= client.Database(os.Getenv("MONGO_DB"))
}
