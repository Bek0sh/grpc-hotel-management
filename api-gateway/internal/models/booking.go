package models

import "github.com/Bek0sh/hotel-management-api-gateway/pkg/pb"

type Booking struct {
	Room         Room   `json:"room"`
	CustomerId   string `json:"customer_id,omitempty"`
	CheckInDate  string `json:"check_in_date,omitempty"`
	CheckOutDate string `json:"check_out_date,omitempty"`
	TotalPrice   int32  `json:"total_price,omitempty"`
}

func (b *Booking) ToProto() *pb.Booking {
	return &pb.Booking{
		CheckInDate:  b.CheckInDate,
		CheckOutDate: b.CheckOutDate,
		CustomerId:   b.CustomerId,
		Room:         b.Room.ToProto(),
	}
}
