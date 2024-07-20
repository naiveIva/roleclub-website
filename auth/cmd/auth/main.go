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
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

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

	service := service.NewService(log, cfg, repo)

	// CreateUsers(service)

	authServer := grpcServer.NewGRPCServer(service)

	grpcServer := grpc.NewServer()

	api.RegisterAuthServer(grpcServer, authServer)

	reflection.Register(grpcServer)

	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", cfg.Server.Port))
	if err != nil {
		log.Error("failed to create listener: %v", logger.Error(err))
		os.Exit(1)
	}

	go func() {
		log.Info("running server", "address", fmt.Sprintf("localhost:%s", cfg.Server.Port))
		err = grpcServer.Serve(l)
		if err != nil {
			log.Error("error while serving: %v", logger.Error(err))
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("stopping server gracefully", slog.String("signal", sign.String()))
	grpcServer.GracefulStop()
	log.Info("application stopped")
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
