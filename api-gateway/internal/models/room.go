package models

import "github.com/Bek0sh/hotel-management-api-gateway/pkg/pb"

type RoomType struct {
	Type  string `json:"type,omitempty"`
	Price int32  `json:"price,omitempty"`
}

func (r *RoomType) ToProto() *pb.RoomType {
	return &pb.RoomType{
		RoomType: r.Type,
		Price:    r.Price,
	}
}

type Room struct {
	RoomNumber int32    `json:"room_number,omitempty"`
	RoomType   RoomType `json:"room_type"`
}

func (r *Room) ToProto() *pb.Room {
	return &pb.Room{
		RoomNumber: r.RoomNumber,
		RoomType:   r.RoomType.ToProto(),
	}
}
