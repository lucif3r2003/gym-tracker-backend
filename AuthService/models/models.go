package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)




type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Password  string             `bson:"password,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Phone     string             `bson:"phone,omitempty"`
	Create_At time.Time          `bson:"create_at,omitempty"`
	Update_At time.Time          `bson:"update_at,omitempty"`
}

type Exercise struct {
    Slug             string   `bson:"slug"`
    Name             string   `bson:"name"`
    PrimaryMuscle    string   `bson:"primary_muscle"`
    SecondaryMuscles []string `bson:"secondary_muscles,omitempty"`
    Category         string   `bson:"category,omitempty"`
    Mechanics        string   `bson:"mechanics,omitempty"`
    Force            string   `bson:"force,omitempty"`
    Difficulty       string   `bson:"difficulty,omitempty"`
}
