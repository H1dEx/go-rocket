package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	orderV1 "github.com/H1dEx/go-rocket/shared/pkg/openapi/order/v1"
	inventoryV1 "github.com/H1dEx/go-rocket/shared/pkg/proto/inventory/v1"
	paymentV1 "github.com/H1dEx/go-rocket/shared/pkg/proto/payment/v1"
)

const (
	httpPort          = "8080"
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second

	inventoryAddress = "localhost:50051"
	paymentAddress   = "localhost:50052"
)

func initInventoryClient() (inventoryV1.InventoryServiceClient, error) {
	conn, err := grpc.NewClient(inventoryAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect inventory: %w", err)
	}

	// defer func() {
	// 	if cerr := conn.Close(); cerr != nil {
	// 		log.Printf("failed to close connect: %v", cerr)
	// 	}
	// }()

	client := inventoryV1.NewInventoryServiceClient(conn)
	return client, nil
}

func initPaymentClient() (paymentV1.PaymentServiceClient, error) {
	conn, err := grpc.NewClient(paymentAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect payment: %w", err)
	}

	// defer func() {
	// 	if cerr := conn.Close(); cerr != nil {
	// 		log.Printf("failed to close connect: %v", cerr)
	// 	}
	// }()

	client := paymentV1.NewPaymentServiceClient(conn)
	return client, nil
}

func main() {
	storage := NewOrderStorage()
	inventoryClient, err := initInventoryClient()
	if err != nil {
		log.Println(err)
		return
	}

	paymentClient, err := initPaymentClient()
	if err != nil {
		log.Println(err)
		return
	}

	orderHandler := NewOrderHandler(storage, paymentClient, inventoryClient)
	orderServer, err := orderV1.NewServer(orderHandler)
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

type OrderHandler struct {
	storage   *OrderStorage
	payment   paymentV1.PaymentServiceClient
	inventory inventoryV1.InventoryServiceClient
}

func NewOrderHandler(storage *OrderStorage, payment paymentV1.PaymentServiceClient, inventory inventoryV1.InventoryServiceClient) *OrderHandler {
	return &OrderHandler{
		storage:   storage,
		payment:   payment,
		inventory: inventory,
	}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *orderV1.CreateOrderRequest) (orderV1.CreateOrderRes, error) {
	orderUUID := uuid.NewString()
	partsUuid := req.PartUuids
	filters := &inventoryV1.PartsFilter{
		Uuids: partsUuid,
	}

	parts, err := h.inventory.ListParts(ctx, &inventoryV1.ListPartsRequest{
		Filter: filters,
	})
	if err != nil {
		return &orderV1.InternalServerError{
			Code:    500,
			Message: fmt.Sprintf("list getting error: %s", err.Error()),
		}, nil
	}

	if len(parts.Parts) < len(partsUuid) {
		return &orderV1.BadRequestError{
			Code:    400,
			Message: "Not all parts finded",
		}, nil
	}

	var totalPrice int
	for _, part := range parts.Parts {
		totalPrice += int(part.Price)
	}
	order := &orderV1.OrderDto{
		OrderUUID:  orderUUID,
		UserUUID:   req.UserUUID,
		PartUuids:  req.PartUuids,
		Status:     orderV1.OrderStatusPENDINGPAYMENT,
		TotalPrice: float32(totalPrice),
	}
	h.storage.mu.Lock()
	defer h.storage.mu.Unlock()
	h.storage.orders[orderUUID] = order

	return &orderV1.CreateOrderResponse{
		OrderUUID:  orderUUID,
		TotalPrice: float32(totalPrice),
	}, nil
}

func (h *OrderHandler) GetOrderByID(ctx context.Context, params orderV1.GetOrderByIDParams) (orderV1.GetOrderByIDRes, error) {
	h.storage.mu.RLock()
	defer h.storage.mu.RUnlock()
	order, ok := h.storage.orders[params.OrderUUID]
	if !ok {
		return &orderV1.NotFoundError{
			Code:    404,
			Message: fmt.Sprintf("sighting with UUID %s not found", params.OrderUUID),
		}, nil
	}

	return &orderV1.GetOrderResponse{
		Order: orderV1.OptOrderDto{Value: *order, Set: true},
	}, nil
}

func (h *OrderHandler) OrderCancelById(ctx context.Context, params orderV1.OrderCancelByIdParams) (orderV1.OrderCancelByIdRes, error) {
	h.storage.mu.Lock()
	defer h.storage.mu.Unlock()
	order, ok := h.storage.orders[params.OrderUUID]

	if !ok {
		return &orderV1.NotFoundError{
			Code:    404,
			Message: fmt.Sprintf("sighting with UUID %s not found", params.OrderUUID),
		}, nil
	}

	if order.Status == orderV1.OrderStatusPAID {
		return &orderV1.ConflictError{
			Code:    409,
			Message: "Order already has been payed",
		}, nil
	}

	h.storage.orders[params.OrderUUID].Status = orderV1.OrderStatusCANCELLED

	return &orderV1.OrderCancelByIdNoContent{}, nil
}

func (h *OrderHandler) PayOrderById(ctx context.Context, req *orderV1.PayOrderRequest, params orderV1.PayOrderByIdParams) (orderV1.PayOrderByIdRes, error) {
	paymentMethod := ToPaymentPaymentMethod(req.PaymentMethod)
	if paymentMethod == paymentV1.PaymentMethod_PAYMENT_METHOD_UNKNOWN || req.PaymentMethod == orderV1.PaymentMethodUNKNOWN {
		return &orderV1.BadRequestError{
			Code:    400,
			Message: "Unknown payment method",
		}, nil
	}

	h.storage.mu.Lock()
	defer h.storage.mu.Unlock()
	part, ok := h.storage.orders[params.OrderUUID]

	if !ok {
		return &orderV1.NotFoundError{
			Code:    404,
			Message: fmt.Sprintf("sighting with UUID %s not found", params.OrderUUID),
		}, nil
	}

	payOrderRes, err := h.payment.PayOrder(ctx, &paymentV1.PayOrderRequest{
		OrderUuid:     params.OrderUUID,
		UserUuid:      part.UserUUID,
		PaymentMethod: paymentMethod,
	})
	if err != nil {
		return &orderV1.InternalServerError{
			Code:    500,
			Message: fmt.Sprintf("pay process error: %s", err.Error()),
		}, nil
	}

	h.storage.orders[params.OrderUUID].PaymentMethod = req.PaymentMethod
	h.storage.orders[params.OrderUUID].TransactionUUID = payOrderRes.TransactionUuid
	h.storage.orders[params.OrderUUID].Status = orderV1.OrderStatusPAID

	return &orderV1.PayOrderResponse{TransactionUUID: payOrderRes.TransactionUuid}, nil
}

func (h *OrderHandler) NewError(_ context.Context, err error) *orderV1.GenericErrorStatusCode {
	return &orderV1.GenericErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: orderV1.GenericError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		},
	}
}

type OrderStorage struct {
	mu     sync.RWMutex
	orders map[string]*orderV1.OrderDto
}

func NewOrderStorage() *OrderStorage {
	return &OrderStorage{
		orders: make(map[string]*orderV1.OrderDto),
	}
}

func ToPaymentPaymentMethod(m orderV1.PaymentMethod) paymentV1.PaymentMethod {
	switch m {
	case orderV1.PaymentMethodCARD:
		return paymentV1.PaymentMethod_PAYMENT_METHOD_CARD
	case orderV1.PaymentMethodSBP:
		return paymentV1.PaymentMethod_PAYMENT_METHOD_SBP
	case orderV1.PaymentMethodCREDITCARD:
		return paymentV1.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD
	case orderV1.PaymentMethodINVESTORMONEY:
		return paymentV1.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY
	default:
		return paymentV1.PaymentMethod_PAYMENT_METHOD_UNKNOWN
	}
}
