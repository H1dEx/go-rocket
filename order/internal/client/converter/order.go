package converter

import (
	"github.com/H1dEx/go-rocket/order/internal/model"
	payment_v1 "github.com/H1dEx/go-rocket/shared/pkg/proto/payment/v1"
)

func OrderPaymentToPaymentMethod(method model.PaymentMethod) payment_v1.PaymentMethod {
	switch method {
	case model.PaymentMethodCard:
		return payment_v1.PaymentMethod_PAYMENT_METHOD_CARD
	case model.PaymentMethodCreditCard:
		return payment_v1.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD
	case model.PaymentMethodInvestorMoney:
		return payment_v1.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY
	case model.PaymentMethodSBP:
		return payment_v1.PaymentMethod_PAYMENT_METHOD_SBP
	default:
		return payment_v1.PaymentMethod_PAYMENT_METHOD_UNKNOWN
	}
}