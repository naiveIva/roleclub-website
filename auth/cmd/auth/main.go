package main

import (
	"auth/internal/config"
	"auth/internal/repository/postgres"
	grpcServer "auth/internal/server"
	"auth/internal/service"
	"auth/models"
	"auth/pkg/database"
	"auth/pkg/logger"
	"context"
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

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Error("failed to open sql connection: %v", logger.Error(err))
	}

	repo := postgres.NewRepository(db)

	service := service.NewService(log, repo)

	// CreateUsers(service)

	authServer := grpcServer.NewGRPCServer(service)

	grpcServer := grpc.NewServer()

	api.RegisterAuthServer(grpcServer, authServer)

	reflection.Register(grpcServer)

	l, err := net.Listen("tcp", cfg.Server.Address)
	if err != nil {
		log.Error("failed to create listener: %v", logger.Error(err))
		os.Exit(1)
	}

	log.Info("running server", "address", cfg.Server.Address)
	err = grpcServer.Serve(l)
	if err != nil {
		log.Error("error while serving: %v", logger.Error(err))
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
		_ = service.RegisterUser(context.TODO(), &user)
	}
}
