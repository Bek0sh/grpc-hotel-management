package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	Id         primitive.ObjectID `bson:"_id" json:"-"`
	BookingId  string             `bson:"booking_id" json:"booking_id"`
	CustomerId *string            `bson:"customer_id" json:"customer_id"`
	Room       struct {
		RoomNumber int `bson:"room_number" json:"room_number"`
		RoomType   struct {
			Type  string `bson:"type" json:"type"`
			Price int    `bson:"price" json:"price"`
		} `bson:"room_type" json:"room_type"`
	} `bson:"room" json:"room"`
	CheckInDate  string `bson:"check_in_date" json:"check_in_date"`
	CheckOutDate string `bson:"check_out_date" json:"check_out_date"`
	TotalPrice   int    `bson:"total_price" json:"total_price"`
}
