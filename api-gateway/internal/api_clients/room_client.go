package api_clients

import (
	"context"
	"github.com/Bek0sh/hotel-management-api-gateway/internal/models"
	"github.com/Bek0sh/hotel-management-api-gateway/pkg/logging"
	"github.com/Bek0sh/hotel-management-api-gateway/pkg/pb"
	"google.golang.org/grpc"
	"time"
)

type RoomClient struct {
	cl  pb.RoomServiceClient
	log *logging.Logger
}

func NewRoomClient(conn *grpc.ClientConn, log *logging.Logger) *RoomClient {
	client := pb.NewRoomServiceClient(conn)
	return &RoomClient{cl: client, log: log}
}

func (r *RoomClient) CreateRoom(room *models.Room) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.CreateRoomRequest{
		Room: room.ToProto(),
	}

	res, err := r.cl.CreateRoom(ctx, req)

	if err != nil {
		r.log.Errorf("failed to create room, error: %v", err)
		return "", err
	}

	return res.GetId(), nil
}

func (r *RoomClient) DeleteRoom(roomNumber int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.DeleteRoomRequest{RoomNumber: roomNumber}

	_, err := r.cl.DeleteRoom(ctx, req)

	if err != nil {
		r.log.Errorf("failed to delete room with roomnumber=%d, error: %v", roomNumber, err)
		return err
	}

	return nil
}

func (r *RoomClient) GetAvailableRooms() ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := r.cl.GetAvailableRooms(ctx, &pb.Empty{})
	if err != nil {
		r.log.Errorf("failed to get all available rooms, error: %v", err)
		return nil, err
	}

	var response []models.Room

	for _, v := range res.GetRooms() {
		room := models.Room{
			RoomNumber: v.RoomNumber,
			RoomType: models.RoomType{
				Type:  v.GetRoomType().GetRoomType(),
				Price: v.GetRoomType().GetPrice(),
			},
		}

		response = append(response, room)
	}
	return response, nil
}

func (r *RoomClient) GetRoomByNumber(roomNumber int32) (*models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.GetRoomByNumberRequest{
		RoomNumber: roomNumber,
	}

	res, err := r.cl.GetRoomByNumber(ctx, req)

	if err != nil {
		r.log.Errorf("failed to get room by number = %d, error: %v", roomNumber, err)
		return nil, err
	}
	room := res.GetRoom()

	response := &models.Room{
		RoomNumber: room.GetRoomNumber(),
		RoomType: models.RoomType{
			Type:  room.GetRoomType().GetRoomType(),
			Price: room.GetRoomType().GetPrice(),
		},
	}

	return response, nil
}

func (r *RoomClient) UpdateRoom(room *models.Room) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.UpdateRoomRequest{
		Room: room.ToProto(),
	}

	_, err := r.cl.UpdateRoom(ctx, req)
	if err != nil {
		r.log.Errorf("failed to update room, error: %v", err)
		return err
	}

	return nil

}
