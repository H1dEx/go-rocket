package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	inventoryApiV1 "github.com/H1dEx/go-rocket/inventory/internal/api/inventory/v1"
	"github.com/H1dEx/go-rocket/inventory/internal/model"
	partRepo "github.com/H1dEx/go-rocket/inventory/internal/repository/part"
	partService "github.com/H1dEx/go-rocket/inventory/internal/service/part"
	inventoryV1 "github.com/H1dEx/go-rocket/shared/pkg/proto/inventory/v1"
	"github.com/joho/godotenv"
)

const grpcPort = 50051

func main() {
	ctx := context.Background()
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("Failed to read env file %v\n", err)
		return
	}

	dbURI := os.Getenv("MONGO_URI")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Printf("failed to connect to database: %v\n", err)
		return
	}
	defer func() {
		cerr := client.Disconnect(ctx)
		if cerr != nil {
			log.Printf("failed to disconnect: %v\n", cerr)
		}
	}()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("failed to ping database: %v\n", err)
		return
	}

	db := client.Database("example")
	storage := partRepo.NewRepository(db)

	id, err := storage.AddPart(ctx, model.Part{
		UUID:          "222",
		Name:          "two",
		Price:         200,
		StockQuantity: 11,
		Category:      "iasdnj",
		CreatedAt:     time.Now(),
	})

	if err != nil {
		log.Printf("error while creating part one: %v", err)
	} else {
		log.Printf("created part one with id: %v", id)
	}
	id, err = storage.AddPart(ctx, model.Part{
		UUID:          "111",
		Name:          "one",
		Price:         100,
		StockQuantity: 11,
		Category:      "bdsnad",
		CreatedAt:     time.Now(),
	})

	if err != nil {
		log.Printf("error while creating part one: %v", err)
	} else {
		log.Printf("created part one with id: %v", id)
	}

	service := partService.NewService(storage)
	api := inventoryApiV1.NewApi(service)

	s := grpc.NewServer()

	inventoryV1.RegisterInventoryServiceServer(s, api)
	reflection.Register(s)

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
