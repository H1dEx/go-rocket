package v1

import (
	"context"

	order_v1 "github.com/H1dEx/go-rocket/shared/pkg/openapi/order/v1"
)

func (a *api) NewError(ctx context.Context, err error) *order_v1.GenericErrorStatusCode {
	return &order_v1.GenericErrorStatusCode{
		StatusCode: 500,
		Response: order_v1.GenericError{
			Code:    500,
			Message: "Internal error",
		},
	}
}
