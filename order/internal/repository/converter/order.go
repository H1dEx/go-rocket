package converter

import (
	"github.com/H1dEx/go-rocket/order/internal/model"
	repoModel "github.com/H1dEx/go-rocket/order/internal/repository/model"
)

func OrderToRepoModel(order model.Order) repoModel.Order {
	return repoModel.Order{
		OrderUUID: order.OrderUUID,
		UserUUID: order.UserUUID,
		PartUuids: order.PartUuids,
		TotalPrice: order.TotalPrice,
		TransactionUUID: order.TransactionUUID,
		PaymentMethod: OrderPaymentMethodToRepoModel(order.PaymentMethod),
		Status: OrderStatusToRepoModel(order.Status),
	}
}

func OrderToModel(order repoModel.Order) model.Order {
	return model.Order{
		OrderUUID: order.OrderUUID,
		UserUUID: order.UserUUID,
		PartUuids: order.PartUuids,
		TotalPrice: order.TotalPrice,
		TransactionUUID: order.TransactionUUID,
		PaymentMethod: OrderPaymentMethodToModel(order.PaymentMethod),
		Status: OrderStatusToModel(order.Status),
	}
}
func OrderPaymentMethodToRepoModel(method model.PaymentMethod) repoModel.PaymentMethod {
	switch method {
	case model.PaymentMethodCard:
		return repoModel.PaymentMethodCard
	case model.PaymentMethodSBP:
		return repoModel.PaymentMethodSBP
	case model.PaymentMethodCreditCard:
		return repoModel.PaymentMethodCreditCard
	case model.PaymentMethodInvestorMoney:
		return repoModel.PaymentMethodInvestorMoney
	default:
		return repoModel.PaymentMethodUnknown
	}
}
func OrderPaymentMethodToModel(method repoModel.PaymentMethod) model.PaymentMethod {
	switch method {
	case repoModel.PaymentMethodCard:
		return model.PaymentMethodCard
	case repoModel.PaymentMethodSBP:
		return model.PaymentMethodSBP
	case repoModel.PaymentMethodCreditCard:
		return model.PaymentMethodCreditCard
	case repoModel.PaymentMethodInvestorMoney:
		return model.PaymentMethodInvestorMoney
	default:
		return model.PaymentMethodUnknown
	}
}

func OrderStatusToRepoModel(status model.OrderStatus) repoModel.OrderStatus {
	switch status {
	case model.OrderStatusPendingPayment:
		return repoModel.OrderStatusPendingPayment
	case model.OrderStatusPaid:
		return repoModel.OrderStatusPaid
	case model.OrderStatusCancelled:
		return repoModel.OrderStatusCancelled
	default:
		return repoModel.OrderStatusUnknown
	}
}

func OrderStatusToModel(status repoModel.OrderStatus) model.OrderStatus {
	switch status {
	case repoModel.OrderStatusPendingPayment:
		return model.OrderStatusPendingPayment
	case repoModel.OrderStatusPaid:
		return model.OrderStatusPaid
	case repoModel.OrderStatusCancelled:
		return model.OrderStatusCancelled
	default:
		return model.OrderStatusUnknown
	}
}

func OrderCreateParamToModel() model.OrderCreateParam {
	return model.OrderCreateParam{
		
	}
}