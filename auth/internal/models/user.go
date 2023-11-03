package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"-" bson:"_id"`
	FullName  string             `json:"full_name" bson:"full_name"`
	Email     string             `json:"email" bson:"email"`
	UserType  string             `json:"user_type" bson:"user_type"`
	Password  string             `json:"-" bson:"password"`
	UserId    string             `json:"user_id" bson:"user_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type RegisterUserRequest struct {
	FullName        string `json:"full_name" bson:"full_name"`
	Email           string `json:"email" bson:"email"`
	UserType        string `json:"user_type" bson:"user_type"`
	Password        string `json:"password" bson:"password"`
	ConfirmPassword string `json:"confirm_password" bson:"confirm_password"`
}

type SignInUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
