package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	orderApi "github.com/H1dEx/go-rocket/order/internal/api/order/v1"
	inventoryCli "github.com/H1dEx/go-rocket/order/internal/client/grpc/inventory/v1"
	paymentCli "github.com/H1dEx/go-rocket/order/internal/client/grpc/payment/v1"
	"github.com/H1dEx/go-rocket/order/internal/migrator"
	orderRepo "github.com/H1dEx/go-rocket/order/internal/repository/order"
	orderService "github.com/H1dEx/go-rocket/order/internal/service/order"
	order_v1 "github.com/H1dEx/go-rocket/shared/pkg/openapi/order/v1"
	inventory_v1 "github.com/H1dEx/go-rocket/shared/pkg/proto/inventory/v1"
	payment_v1 "github.com/H1dEx/go-rocket/shared/pkg/proto/payment/v1"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	httpPort          = "8080"
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second

	inventoryAddress = "localhost:50051"
	paymentAddress   = "localhost:50052"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("Failed to read env file %v\n", err)
		return
	}

	dbURI := os.Getenv("DB_URI")
	pool, err := pgxpool.New(ctx, dbURI)
	if err != nil {
		log.Printf("failed to connect to database: %v\n", err)
		return
	}

	defer pool.Close()

	migrationsDir := os.Getenv("MIGRATIONS_DIR")
	migratorRunner := migrator.NewMigrator(stdlib.OpenDB(*pool.Config().ConnConfig.Copy()), migrationsDir)

	if err := migratorRunner.Up(); err != nil {
		log.Printf("Ошибка миграции базы данных: %v\n", err)
		return
	}

	paymentConn, err := grpc.NewClient(paymentAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect payment: %s", err.Error())
		return
	}
	defer func() {
		if cerr := paymentConn.Close(); cerr != nil {
			log.Printf("failed to close connect: %v", cerr)
		}
	}()
	paymentClient := paymentCli.NewClient(payment_v1.NewPaymentServiceClient(paymentConn))

	inventoryConn, err := grpc.NewClient(inventoryAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect inventory: %s", err.Error())
		return
	}

	defer func() {
		if cerr := inventoryConn.Close(); cerr != nil {
			log.Printf("failed to close connect: %v", cerr)
		}
	}()

	inventoryClient := inventoryCli.NewClient(inventory_v1.NewInventoryServiceClient(inventoryConn))

	storage := orderRepo.NewRepository(pool)
	service := orderService.NewService(storage, inventoryClient, paymentClient)
	api := orderApi.NewApi(service)
	orderServer, err := order_v1.NewServer(api)
	if err != nil {
		log.Fatalf("ошибка создания сервера OpenAPI: %v", err)
		return
	}

	r := chi.NewRouter()

	// Добавляем middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/", orderServer)

	server := &http.Server{
		Addr:              net.JoinHostPort("localhost", httpPort),
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	go func() {
		log.Printf("🚀 HTTP-сервер запущен на порту %s\n", httpPort)
		err = server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("❌ Ошибка запуска сервера: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Завершение работы сервера...")

	// Создаем контекст с таймаутом для остановки сервера
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Printf("❌ Ошибка при остановке сервера: %v\n", err)
	}

	log.Println("✅ Сервер остановлен")
}
