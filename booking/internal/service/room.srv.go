package service

import (
	"context"
	"fmt"
	"github.com/Bek0sh/hotel-management-booking/internal/models"
	"github.com/Bek0sh/hotel-management-booking/pkg/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *roomService) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	id := primitive.NewObjectID()
	roomId := id.Hex()
	ok := r.repo.RoomType.CategoryExists(ctx, req.GetRoom().GetRoomTypeId())
	if !ok {
		return nil, fmt.Errorf("failed to find room type that you provided, room type=%s", req.GetRoom().GetRoomType())
	}
	createRoom := models.Room{
		Id:          id,
		RoomId:      roomId,
		RoomNumber:  int(req.GetRoom().GetRoomNumber()),
		RoomTypeId:  req.GetRoom().GetRoomTypeId(),
		IsAvailable: true,
	}

	err := r.repo.Room.CreateRoom(ctx, &createRoom)
	if err != nil {
		return nil, err
	}

	return &pb.CreateRoomResponse{Id: roomId}, nil
}

func (r *roomService) DeleteRoom(ctx context.Context, req *pb.DeleteRoomRequest) (*pb.Empty, error) {
	err := r.repo.Room.DeleteRoomByNumber(ctx, int(req.GetRoomNumber()))

	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
func (r *roomService) GetAvailableRooms(ctx context.Context, _ *pb.Empty) (*pb.GetAvailableRoomsResponse, error) {
	var response []*pb.Room

	res, err := r.repo.Room.GetAllAvailableRooms(ctx)

	if err != nil {
		return nil, err
	}

	for _, v := range res {

		roomType, err := r.repo.RoomType.GetRoomTypeByType(ctx, v.RoomTypeId)
		if err != nil {
			return nil, err
		}

		room := &pb.Room{
			RoomNumber: int32(v.RoomNumber),
			RoomType: &pb.RoomType{
				RoomType: roomType.Type,
				Price:    int32(roomType.Price),
			},
		}
		response = append(response, room)
	}

	return &pb.GetAvailableRoomsResponse{Rooms: response}, nil
}
func (r *roomService) UpdateRoom(ctx context.Context, req *pb.UpdateRoomRequest) (*pb.Empty, error) {
	room := models.Room{
		RoomNumber: int(req.GetRoom().GetRoomNumber()),
		RoomTypeId: req.GetRoom().GetRoomTypeId(),
	}

	roomType, err := r.repo.RoomType.GetRoomTypeByType(ctx, req.GetRoom().GetRoomTypeId())
	if err != nil {
		return nil, err
	}

	room.RoomType = *roomType
	err = r.repo.Room.UpdateRoom(ctx, &room)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
func (r *roomService) GetRoomByNumber(ctx context.Context, req *pb.GetRoomByNumberRequest) (*pb.GetRoomByNumberResponse, error) {
	res, err := r.repo.Room.GetRoomByNumber(ctx, int(req.GetRoomNumber()))
	if err != nil {
		return nil, err
	}

	roomType, err := r.repo.RoomType.GetRoomTypeByType(ctx, res.RoomTypeId)
	if err != nil {
		return nil, err
	}
	response := &pb.Room{
		RoomNumber: int32(res.RoomNumber),
		RoomType: &pb.RoomType{
			RoomType: roomType.Type,
			Price:    int32(roomType.Price),
		},
	}

	return &pb.GetRoomByNumberResponse{Room: response}, nil
}
