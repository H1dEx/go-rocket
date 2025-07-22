package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	inventoryApiV1 "github.com/H1dEx/go-rocket/inventory/internal/api/inventory/v1"
	partRepo "github.com/H1dEx/go-rocket/inventory/internal/repository/part"
	partService "github.com/H1dEx/go-rocket/inventory/internal/service/part"
	inventoryV1 "github.com/H1dEx/go-rocket/shared/pkg/proto/inventory/v1"
)

const grpcPort = 50051

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return
	}
	defer func() {
		if cerr := listener.Close(); cerr != nil {
			log.Printf("failed to close listener: %v\n", cerr)
		}
	}()
	s := grpc.NewServer()

	storage := partRepo.NewRepository()

	service := partService.NewService(storage)
	api := inventoryApiV1.NewApi(service)

	inventoryV1.RegisterInventoryServiceServer(s, api)
	reflection.Register(s)

	go func() {
		log.Printf("🚀 gRPC server listening on %d\n", grpcPort)
		if err := s.Serve(listener); err != nil && !errors.Is(err, net.ErrClosed) {
			log.Printf("failed to serve: %v\n", err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down gRPC server...")
	s.GracefulStop()
	log.Println("✅ Server stopped")
}
