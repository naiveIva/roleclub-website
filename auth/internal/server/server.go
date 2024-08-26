package grpcserver

import (
	"context"
	"errors"
	"fmt"

	api "auth/api/gen"
	"auth/internal/service"
	"auth/models"
	"auth/pkg/autherrors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCServer struct {
	api.UnimplementedAuthServer
	service *service.Service
}

func NewGRPCServer(s *service.Service) *GRPCServer {
	return &GRPCServer{
		service: s,
	}
}

func (s *GRPCServer) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	switch {
	case req.GetFirstName() == "":
		return nil, status.Error(codes.InvalidArgument, "first name is required")
	case req.GetLastName() == "":
		return nil, status.Error(codes.InvalidArgument, "last name is required")
	case req.GetFatherName() == "":
		return nil, status.Error(codes.InvalidArgument, "father name is required")
	case req.GetTelNumber() == "":
		return nil, status.Error(codes.InvalidArgument, "telephone number is required")
	case req.GetPassword() == "":
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	err := s.service.RegisterUser(
		ctx, &models.User{
			FirstName:    req.GetFirstName(),
			LastName:     req.GetLastName(),
			FatherName:   req.GetFatherName(),
			TelNumber:    req.GetTelNumber(),
			Password:     req.GetPassword(),
			IsHSEStudent: req.GetIsHSEStudent(),
		},
	)

	if errors.Is(err, autherrors.ErrUserAlreadyExists) {
		return nil, status.Error(codes.AlreadyExists, "user already exists")
	}

	return &api.CreateUserResponse{Ok: "all done"}, nil
}

func (s *GRPCServer) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	switch {
	case req.GetTelNumber() == "":
		return nil, status.Error(codes.InvalidArgument, "telephone number is required")
	case req.GetPassword() == "":
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	token, err := s.service.Login(ctx, req.GetTelNumber(), req.GetPassword())
	if err != nil {
		if errors.Is(err, autherrors.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		if errors.Is(err, autherrors.ErrWrongPassword) {
			return nil, status.Error((codes.Unauthenticated), "wrong password")
		}
		return nil, err
	}
	str := fmt.Sprintf("jwt token: %s", token)
	return &api.LoginResponse{Token: str}, nil
}
