package service

import (
	"context"
	"github.com/Bek0sh/hotel-management-booking/internal/models"
	"github.com/Bek0sh/hotel-management-booking/pkg/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *roomService) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	id := primitive.NewObjectID()
	roomId := id.Hex()
	createRoom := models.Room{
		Id:         id,
		RoomId:     roomId,
		RoomNumber: int(req.GetRoom().GetRoomNumber()),
		RoomType: struct {
			Type  string `bson:"type" json:"type"`
			Price int    `bson:"price" json:"price"`
		}{
			Type:  req.GetRoom().GetRoomType().GetRoomType(),
			Price: int(req.GetRoom().GetRoomType().GetPrice()),
		},
		IsAvailable: true,
	}

	err := r.repo.CreateRoom(ctx, &createRoom)
	if err != nil {
		return nil, err
	}

	return &pb.CreateRoomResponse{Id: roomId}, nil
}

func (r *roomService) DeleteRoom(ctx context.Context, req *pb.DeleteRoomRequest) (*pb.Empty, error) {
	err := r.repo.DeleteRoomByNumber(ctx, int(req.GetRoomNumber()))

	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
func (r *roomService) GetAvailableRooms(ctx context.Context, _ *pb.Empty) (*pb.GetAvailableRoomsResponse, error) {
	var response []*pb.Room

	res, err := r.repo.GetAllAvailableRooms(ctx)

	if err != nil {
		return nil, err
	}

	for _, v := range res {
		room := &pb.Room{
			RoomNumber: int32(v.RoomNumber),
			RoomType: &pb.RoomType{
				RoomType: v.RoomType.Type,
				Price:    int32(v.RoomType.Price),
			},
		}
		response = append(response, room)
	}

	return &pb.GetAvailableRoomsResponse{Rooms: response}, nil
}
func (r *roomService) UpdateRoom(ctx context.Context, req *pb.UpdateRoomRequest) (*pb.Empty, error) {
	room := models.Room{
		RoomNumber: int(req.GetRoom().GetRoomNumber()),
		RoomType: struct {
			Type  string `bson:"type" json:"type"`
			Price int    `bson:"price" json:"price"`
		}{
			Type:  req.GetRoom().GetRoomType().GetRoomType(),
			Price: int(req.GetRoom().GetRoomType().GetPrice())},
	}

	err := r.repo.UpdateRoom(ctx, &room, false)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
func (r *roomService) GetRoomByNumber(ctx context.Context, req *pb.GetRoomByNumberRequest) (*pb.GetRoomByNumberResponse, error) {
	res, err := r.repo.GetRoomByNumber(ctx, int(req.GetRoomNumber()))
	if err != nil {
		return nil, err
	}

	response := &pb.Room{
		RoomNumber: int32(res.RoomNumber),
		RoomType: &pb.RoomType{
			RoomType: res.RoomType.Type,
			Price:    int32(res.RoomType.Price),
		},
	}

	return &pb.GetRoomByNumberResponse{Room: response}, nil
}
