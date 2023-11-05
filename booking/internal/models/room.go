package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Room struct {
	Id          primitive.ObjectID `bson:"_id" json:"-"`
	RoomNumber  int                `bson:"room_number" json:"room_number"`
	IsAvailable bool               `bson:"is_available" json:"-"`
	RoomId      string             `bson:"room_id" json:"-"`
	RoomTypeId  string             `bson:"room_type_id" json:"-"`
	RoomType    RoomType           `bson:"-" json:"room_type"`
}

type RoomType struct {
	Id         primitive.ObjectID `json:"-,omitempty" bson:"_id"`
	Type       string             `json:"type,omitempty" bson:"type"`
	Price      int                `json:"price,omitempty" bson:"price"`
	RoomTypeId string             `json:"room_type_id,omitempty" bson:"room_type_id"`
}
