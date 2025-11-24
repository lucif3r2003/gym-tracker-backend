package repositories

import (
	"context"
	"fmt"
	"gym-tracker-project/database"
	"gym-tracker-project/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)


func UserCollection() *mongo.Collection{
	return database.DB.Collection("users")
}

func CreateUser(user models.User) error{
	_, err := UserCollection().InsertOne(context.Background(), user)
	return err
}

func UpdateUser(user models.User) error{
	id, _ := primitive.ObjectIDFromHex(user.ID.String())

	user.Update_At = time.Now()
	update := bson.M{"$set": user}
	result, err := UserCollection().UpdateOne(context.TODO(), bson.M{"_id":id}, update)
	fmt.Printf("\nupdate %v documents !", result.ModifiedCount)
	return err
}
//-------------------------------------------------------------------
func FindUserByEmail(mail string) (models.User, error){
	var user models.User
	err := UserCollection().FindOne(context.TODO(), bson.M{"email" : mail}).Decode(&user)
	return user, err
}



