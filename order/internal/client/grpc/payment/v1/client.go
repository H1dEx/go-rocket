package v1

import (
	def "github.com/H1dEx/go-rocket/order/internal/client/grpc"
	payment_v1 "github.com/H1dEx/go-rocket/shared/pkg/proto/payment/v1"
)

var _ def.PaymentClient = (*client)(nil)


type client struct {
	genCli payment_v1.PaymentServiceClient
}

func NewClient(cli payment_v1.PaymentServiceClient) *client {
	return &client{
		genCli: cli,
	}
}