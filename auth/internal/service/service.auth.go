package service

import (
	"context"
	"fmt"
	"github.com/Bek0sh/hotel-management-auth/pkg/pb"
	"strconv"
	"time"

	"github.com/Bek0sh/hotel-management-auth/internal/models"
	"github.com/Bek0sh/hotel-management-auth/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *service) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	if violations := validateSignIn(req); violations != nil {
		s.logger.Error("failed to validate sign in user request")
		return nil, InvalidArgumentError(violations)
	}

	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		s.logger.Errorf("failed to grab user with email=%s, error: %v", req.Email, err)
		return nil, err
	}

	if err := utils.ComparePassword(req.Password, user.Password); err != nil {
		s.logger.Error("input password and database password does not match")
		return nil, err
	}

	duration, _ := strconv.Atoi(s.cfg.JWT.AccessTokenDur)

	accessToken, err := s.jwt.CreateToken(user.UserId, user.UserType, time.Duration(duration*int(time.Minute)))
	if err != nil {
		s.logger.Error("failed to create access token")
		return nil, err
	}

	response := &pb.SignInResponse{AccessToken: accessToken}
	return response, nil
}

func (s *service) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	if violations := validateRegisterUser(req); violations != nil {
		s.logger.Error("failed to validate register user request")
		return nil, InvalidArgumentError(violations)
	}

	if req.ConfirmPassword != req.Password {
		s.logger.Error("passwords do not match")
		return nil, fmt.Errorf("confirm password and password must be equal")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password, error: %v", err)
	}

	id := primitive.NewObjectID()
	userId := id.Hex()

	userToCreate := models.User{
		Id:        id,
		FullName:  req.FullName,
		Email:     req.Email,
		Password:  hashedPassword,
		UserId:    userId,
		UserType:  req.UserType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.CreateUser(ctx, &userToCreate); err != nil {
		s.logger.Error("register user in service failed")
		return nil, err
	}

	response := &pb.RegisterResponse{Id: userId}

	return response, nil
}

func (s *service) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	payload, err := s.jwt.VerifyToken(req.GetAccessToken())
	if err != nil {
		s.logger.Errorf("failed to validate token, error: %v", err)
		return nil, err
	}

	response := &pb.ValidateTokenResponse{Id: payload.Id, UserType: payload.UserRole}

	return response, nil
}
func (s *service) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	user, err := s.repo.GetUserById(ctx, req.GetId())
	if err != nil {
		s.logger.Errorf("failed to grab user with id=%s, error: %v", req.GetId(), err)
		return nil, err
	}

	response := &pb.GetResponse{
		User: &pb.User{
			Id:        user.UserId,
			FullName:  user.FullName,
			Email:     user.Email,
			UserType:  user.UserType,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
	}
	return response, nil
}
