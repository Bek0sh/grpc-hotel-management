package api_clients

import (
	"context"
	"github.com/Bek0sh/hotel-management-api-gateway/internal/models"
	"github.com/Bek0sh/hotel-management-api-gateway/pkg/logging"
	"github.com/Bek0sh/hotel-management-api-gateway/pkg/pb"
	"google.golang.org/grpc"
	"time"
)

type RoomTypeClient struct {
	cl  pb.RoomTypeServiceClient
	log *logging.Logger
}

func NewRoomTypeClient(conn *grpc.ClientConn, log *logging.Logger) *RoomTypeClient {
	client := pb.NewRoomTypeServiceClient(conn)
	return &RoomTypeClient{log: log, cl: client}
}

func (r *RoomTypeClient) CreateRoomType(roomType *models.RoomType) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.CreateRoomTypeReq{
		Type:  roomType.Type,
		Price: roomType.Price,
	}

	res, err := r.cl.CreateRoomType(ctx, req)

	if err != nil {
		r.log.Errorf("failed to create room type, error: %v", err)
		return "", err
	}

	return res.GetId(), nil
}

func (r *RoomTypeClient) GetRoomTypeByType(typ string) (*models.RoomType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.GetByRoomTypeReq{RoomType: typ}

	res, err := r.cl.GetByRoomType(ctx, req)
	if err != nil {
		r.log.Errorf("failed to get room type by type, type=%s, error: %v", typ, err)
		return nil, err
	}

	response := &models.RoomType{
		Type:  res.GetRoomType().GetRoomType(),
		Price: res.GetRoomType().GetPrice(),
	}

	return response, nil
}

func (r *RoomTypeClient) GetAllRoomTypes() ([]models.RoomType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := r.cl.GetAllRoomTypes(ctx, &pb.Empty{})
	if err != nil {
		r.log.Errorf("failed to get all room types, error: %v", err)
		return nil, err
	}

	var response []models.RoomType

	for _, v := range res.GetRoomType() {
		roomType := models.RoomType{
			Type:  v.GetRoomType(),
			Price: v.GetPrice(),
		}

		response = append(response, roomType)
	}

	return response, nil
}

func (r *RoomTypeClient) UpdateRoomType(roomType *models.RoomType) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.UpdateRoomTypeReq{
		RoomType: roomType.ToProto(),
	}

	_, err := r.cl.UpdateRoomType(ctx, req)

	if err != nil {
		r.log.Errorf("failed to update room type, error: %v", err)
		return err
	}

	return nil
}

func (r *RoomTypeClient) DeleteRoomType(typ string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.DeleteRoomTypeReq{Type: typ}

	_, err := r.cl.DeleteRoomType(ctx, req)
	if err != nil {
		r.log.Errorf("failed to delete room type with type=%s, error: %v", typ, err)
		return err
	}

	return nil
}
