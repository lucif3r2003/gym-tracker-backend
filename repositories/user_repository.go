package repositories

import (
	"context"
	"fmt"
	"gym-tracker-project/database"
	"gym-tracker-project/models"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
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

func DeleteUser(mail string) error{
	_, err:= FindUserByEmail(mail)
	if err != nil{
		log.Fatal(err)
	}
	result, err:= UserCollection().DeleteOne(context.Background(), bson.M{"email": mail})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("delete %v documents !", result.DeletedCount)
	return err
}

//-------------------------------------------------------------------
func FindUserByEmail(mail string) (models.User, error){
	var user models.User
	err := UserCollection().FindOne(context.TODO(), bson.M{"email" : mail}).Decode(&user)
	return user, err
}

func CheckDuplicateEmail(email string) (bool, error) {
	count, err := UserCollection().CountDocuments(context.TODO(), bson.M{"email": email})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}


func CheckPassword(hashed string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
