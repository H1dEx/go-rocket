package converter

import (
	"github.com/H1dEx/go-rocket/order/internal/model"
	order_v1 "github.com/H1dEx/go-rocket/shared/pkg/openapi/order/v1"
)

func OrderStatusToDtoModel(status model.OrderStatus) order_v1.OrderStatus {
	switch status {
	case model.OrderStatusCancelled:
		return order_v1.OrderStatusCANCELLED
	case model.OrderStatusPaid:
		return order_v1.OrderStatusPAID
	case model.OrderStatusPendingPayment:
		return order_v1.OrderStatusPENDINGPAYMENT
	default:
		return order_v1.OrderStatusUNKNOWN
	}
}

func OrderStatusToModel(status order_v1.OrderStatus) model.OrderStatus {
	switch status {
	case order_v1.OrderStatusCANCELLED:
		return model.OrderStatusCancelled
	case order_v1.OrderStatusPAID:
		return model.OrderStatusPaid
	case order_v1.OrderStatusPENDINGPAYMENT:
		return model.OrderStatusPendingPayment
	default:
		return model.OrderStatusUnknown
	}
}

func OrderPaymentToDtoModel(payment model.PaymentMethod) order_v1.PaymentMethod {
	switch payment {
	case model.PaymentMethodCard:
		return order_v1.PaymentMethodCARD
	case model.PaymentMethodSBP:
		return order_v1.PaymentMethodSBP
	case model.PaymentMethodCreditCard:
		return order_v1.PaymentMethodCREDITCARD
	case model.PaymentMethodInvestorMoney:
		return order_v1.PaymentMethodINVESTORMONEY
	default:
		return order_v1.PaymentMethodUNKNOWN
	}
}
func OrderPaymentToModel(payment order_v1.PaymentMethod) model.PaymentMethod {
	switch payment {
	case order_v1.PaymentMethodCARD:
		return model.PaymentMethodCard
	case order_v1.PaymentMethodSBP:
		return model.PaymentMethodSBP
	case order_v1.PaymentMethodCREDITCARD:
		return model.PaymentMethodCreditCard
	case order_v1.PaymentMethodINVESTORMONEY:
		return model.PaymentMethodInvestorMoney
	default:
		return model.PaymentMethodUnknown
	}
}

func OrderToDtoModel(order model.Order) order_v1.OrderDto {
	return order_v1.OrderDto{
		OrderUUID:       order.OrderUUID,
		UserUUID:        order.UserUUID,
		PartUuids:       order.PartUuids,
		TotalPrice:      order.TotalPrice,
		TransactionUUID: order.TransactionUUID,
		PaymentMethod:   OrderPaymentToDtoModel(order.PaymentMethod),
		Status:          OrderStatusToDtoModel(order.Status),
	}
}

func OrderToModel(order order_v1.OrderDto) model.Order {
	return model.Order{
		OrderUUID:       order.OrderUUID,
		UserUUID:        order.UserUUID,
		PartUuids:       order.PartUuids,
		TotalPrice:      order.TotalPrice,
		TransactionUUID: order.TransactionUUID,
		PaymentMethod: OrderPaymentToModel(order.PaymentMethod),
		Status: OrderStatusToModel(order.Status),
	}
}