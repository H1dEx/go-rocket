package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"slices"
	"sync"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	inventoryV1 "github.com/H1dEx/go-rocket/shared/pkg/proto/inventory/v1"
)

const grpcPort = 50051

type InventoryStorage struct {
	mu    sync.RWMutex
	parts map[string]*inventoryV1.Part
}

func NewInventoryStorage() *InventoryStorage {
	return &InventoryStorage{
		parts: make(map[string]*inventoryV1.Part),
	}
}

func (s *InventoryStorage) GetPart(partUUID string) *inventoryV1.Part {
	s.mu.RLock()
	defer s.mu.RUnlock()

	part, ok := s.parts[partUUID]

	if !ok {
		return nil
	}

	return part
}

func validatePart(p *inventoryV1.Part, filters *inventoryV1.PartsFilter) bool {
	if len(filters.Uuids) > 0 && !slices.Contains(filters.Uuids, p.Uuid) {
		return false
	}

	if len(filters.Names) > 0 && !slices.Contains(filters.Names, p.Name) {
		return false
	}

	if len(filters.Categories) > 0 && !slices.Contains(filters.Categories, p.Category) {
		return false
	}

	if len(filters.ManufacturerCountries) > 0 && !slices.Contains(filters.ManufacturerCountries, p.Manufacturer.Country) {
		return false
	}

	if len(filters.Tags) > 0 {
		match := false
		for _, tag := range filters.Tags {
			if slices.Contains(p.Tags, tag) {
				match = true
				break
			}
		}
		if !match {
			return false
		}
	}

	return true
}

func (s *InventoryStorage) GetListParts(filters *inventoryV1.PartsFilter) []*inventoryV1.Part {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var filteredParts []*inventoryV1.Part
	for _, part := range s.parts {
		if validatePart(part, filters) {
			filteredParts = append(filteredParts, part)
		}
	}

	return filteredParts
}

type inventoryService struct {
	inventoryV1.UnimplementedInventoryServiceServer
	storage *InventoryStorage
}

func (s *inventoryService) GetPart(ctx context.Context, in *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error) {
	uuid := in.GetUuid()
	if uuid == "" {
		return nil, status.Error(codes.InvalidArgument, "UUID is empty")
	}
	part := s.storage.GetPart(uuid)

	if part == nil {
		return nil, status.Errorf(codes.NotFound, "part with UUID %s not found", uuid)
	}

	return &inventoryV1.GetPartResponse{
		Part: part,
	}, nil
}

func (s *inventoryService) ListParts(ctx context.Context, in *inventoryV1.ListPartsRequest) (*inventoryV1.ListPartsResponse, error) {
	if in.Filter == nil {
		return nil, status.Error(codes.InvalidArgument, "Filters is invalid")
	}
	parts := s.storage.GetListParts(in.Filter)

	if len(parts) == 0 {
		return nil, status.Error(codes.NotFound, "Not found")
	}

	return &inventoryV1.ListPartsResponse{Parts: parts}, nil
}

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
	storage := NewInventoryStorage()
	firstMock := &inventoryV1.Part{
		Uuid:  "111",
		Name:  "First detail",
		Price: 100,
	}
	secondMock := &inventoryV1.Part{
		Uuid:  "222",
		Name:  "Second detail",
		Price: 200,
	}
	storage.parts[firstMock.Uuid] = firstMock
	storage.parts[secondMock.Uuid] = secondMock
	service := &inventoryService{
		storage: storage,
	}

	inventoryV1.RegisterInventoryServiceServer(s, service)
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
