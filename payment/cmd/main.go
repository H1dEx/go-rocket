package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	paymentApiV1 "github.com/H1dEx/go-rocket/payment/internal/api/payment/v1"
	paymentService "github.com/H1dEx/go-rocket/payment/internal/service/payment"
	paymentV1 "github.com/H1dEx/go-rocket/shared/pkg/proto/payment/v1"
)

const grpcPort = 50052

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

	service := paymentService.NewService()
	api := paymentApiV1.NewApi(service)

	paymentV1.RegisterPaymentServiceServer(s, api)
	reflection.Register(s)
	go func() {
		log.Printf("🚀 gRPC server listening on %d\n", grpcPort)
		if err := s.Serve(listener); err != nil {
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
