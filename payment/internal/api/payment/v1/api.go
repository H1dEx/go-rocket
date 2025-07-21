package v1

import (
	"github.com/H1dEx/go-rocket/payment/internal/service"
	payment_v1 "github.com/H1dEx/go-rocket/shared/pkg/proto/payment/v1"
)

type api struct {
	payment_v1.UnimplementedPaymentServiceServer
	service service.PaymentService
}

func NewApi(service service.PaymentService) *api {
	return &api{
		service: service,
	}
}