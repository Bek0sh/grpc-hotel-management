package service

import (
	"context"
	"github.com/Bek0sh/hotel-management-booking/internal/models"
	"github.com/Bek0sh/hotel-management-booking/pkg/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *roomTypeService) CreateRoomType(ctx context.Context, req *pb.CreateRoomTypeReq) (*pb.CreateRoomTypeRes, error) {
	id := primitive.NewObjectID()
	roomTypeId := id.Hex()

	rt := &models.RoomType{
		Id:         id,
		RoomTypeId: roomTypeId,
		Type:       req.GetType(),
		Price:      int(req.GetPrice()),
	}

	err := r.repo.CreateRoomType(ctx, rt)
	if err != nil {
		return nil, err
	}

	return &pb.CreateRoomTypeRes{Id: roomTypeId}, nil
}
func (r *roomTypeService) DeleteRoomType(ctx context.Context, req *pb.DeleteRoomTypeReq) (*pb.Empty, error) {
	err := r.repo.DeleteRoomTypeByType(ctx, req.Type)

	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
func (r *roomTypeService) UpdateRoomType(ctx context.Context, req *pb.UpdateRoomTypeReq) (*pb.Empty, error) {
	err := r.repo.UpdateRoomType(ctx, req.GetRoomType().GetRoomType(), int(req.GetRoomType().GetPrice()))

	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (r *roomTypeService) GetAllRoomTypes(ctx context.Context, req *pb.Empty) (*pb.GetAllRoomTypesRes, error) {
	res, err := r.repo.GetAllRoomTypes(ctx)

	if err != nil {
		return nil, err
	}

	var response []*pb.RoomType

	for _, v := range res {
		ans := &pb.RoomType{
			RoomType: v.Type,
			Price:    int32(v.Price),
		}
		response = append(response, ans)
	}

	return &pb.GetAllRoomTypesRes{RoomType: response}, nil
}
func (r *roomTypeService) GetByRoomType(ctx context.Context, req *pb.GetByRoomTypeReq) (*pb.GetByRoomTypeRes, error) {
	res, err := r.repo.GetRoomTypeByType(ctx, req.GetRoomType())

	if err != nil {
		return nil, err
	}

	response := &pb.RoomType{
		RoomType: res.Type,
		Price:    int32(res.Price),
	}

	return &pb.GetByRoomTypeRes{RoomType: response}, nil
}
