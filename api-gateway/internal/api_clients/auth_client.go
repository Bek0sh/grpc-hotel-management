package api_clients

import (
	"context"
	"github.com/Bek0sh/hotel-management-api-gateway/internal/models"
	"github.com/Bek0sh/hotel-management-api-gateway/pkg/logging"
	"github.com/Bek0sh/hotel-management-api-gateway/pkg/pb"
	"google.golang.org/grpc"
	"time"
)

type AuthClient struct {
	cl  pb.AuthServiceClient
	log *logging.Logger
}

func NewAuthClient(conn *grpc.ClientConn, log *logging.Logger) *AuthClient {
	client := pb.NewAuthServiceClient(conn)
	return &AuthClient{log: log, cl: client}
}

func (a *AuthClient) RegisterUser(user *models.RegisterUser) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	req := &pb.RegisterRequest{
		FullName:        user.FullName,
		Email:           user.Email,
		UserType:        user.UserType,
		Password:        user.Password,
		ConfirmPassword: user.ConfirmPassword,
	}
	res, err := a.cl.Register(ctx, req)
	if err != nil {
		a.log.Errorf("failed to register user, error: %v", err)
		return "", err
	}

	return res.GetId(), nil
}

func (a *AuthClient) SignInUser(in *models.SignIn) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	req := &pb.SignInRequest{
		Password: in.Password,
		Email:    in.Email,
	}
	res, err := a.cl.SignIn(ctx, req)

	if err != nil {
		a.log.Errorf("failed to sign in user with email=%s, error: %v", in.Email, err)
		return "", err
	}

	return res.GetAccessToken(), nil
}

func (a *AuthClient) ValidateAccessToken(token string) (string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	req := &pb.ValidateTokenRequest{AccessToken: token}
	res, err := a.cl.ValidateToken(ctx, req)

	if err != nil {
		a.log.Errorf("failed to validate token, error: %v", err)
		return "", "", err
	}

	return res.GetId(), res.GetUserType(), nil
}

func (a *AuthClient) GetUserProfile(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.GetUserByIdRequest{Id: id}

	res, err := a.cl.GetUserById(ctx, req)

	if err != nil {
		a.log.Errorf("failed to find user with id=%s, error: %v", id, err)
		return nil, err
	}
	user := res.GetUser()

	response := &models.User{
		FullName:  user.GetFullName(),
		UserId:    user.GetId(),
		Email:     user.GetEmail(),
		UserType:  user.GetUserType(),
		CreatedAt: user.GetCreatedAt(),
		UpdatedAt: user.GetUpdatedAt(),
	}
	return response, nil
}
