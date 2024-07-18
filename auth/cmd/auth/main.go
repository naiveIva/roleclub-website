package main

import (
	"auth/internal/config"
	"auth/internal/logger"
	"auth/internal/repository/postgres"
	grpcserver "auth/internal/server"
	"auth/internal/service"
	"auth/models"
	"flag"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	api "auth/api/gen"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "config/config.yaml", "path to config file")
}

func main() {
	flag.Parse()

	cfg := config.MustInit(configPath)

	log := logger.Init()

	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		log.Error("failed to open sql connection: %v", err)
	}

	repo := postgres.NewRepository(db)

	service := service.NewService(repo)

	// CreateUsers(service)

	AuthServer := grpcserver.NewGRPCServer(service)

	GrpcServer := grpc.NewServer()

	api.RegisterAuthServer(GrpcServer, AuthServer)

	reflection.Register(GrpcServer)

	l, err := net.Listen("tcp", "localhost:8082")
	if err != nil {
		log.Error("failed to create listener: %v", err)
		os.Exit(1)
	}

	err = GrpcServer.Serve(l)
	if err != nil {
		log.Error("error while serving: %v", err)
	}
}

func CreateUsers(service *service.Service) {
	users := []models.User{
		{
			FirstName:    "Иоанна",
			LastName:     "Мысниченко",
			FatherName:   "Николаевна",
			TelNumber:    "12345",
			Password:     "qwerty",
			IsHSEStudent: false,
		},
		{
			FirstName:    "Ангелина",
			LastName:     "Никитич",
			FatherName:   "Батьковна",
			TelNumber:    "67890",
			Password:     "ффффффф",
			IsHSEStudent: true,
		},
		{
			FirstName:    "Ксения",
			LastName:     "Борщева",
			FatherName:   "Батьковна",
			TelNumber:    "000",
			Password:     "fffjjjjfjfjf",
			IsHSEStudent: false,
		},
	}

	for _, user := range users {
		_ = service.CreateUser(&user)
	}
}
