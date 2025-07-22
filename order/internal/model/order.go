package model

type PaymentMethod string

const (
	PaymentMethodUnknown       PaymentMethod = "UNKNOWN"
	PaymentMethodCard          PaymentMethod = "CARD"
	PaymentMethodSBP           PaymentMethod = "SBP"
	PaymentMethodCreditCard    PaymentMethod = "CREDIT_CARD"
	PaymentMethodInvestorMoney PaymentMethod = "INVESTOR_MONEY"
)

type OrderStatus string

const (
	OrderStatusUnknown        OrderStatus = "UNKNOWN"
	OrderStatusPendingPayment OrderStatus = "PENDING_PAYMENT"
	OrderStatusPaid           OrderStatus = "PAID"
	OrderStatusCancelled      OrderStatus = "CANCELLED"
)

type Order struct {
	OrderUUID       string
	UserUUID        string
	PartUuids       []string
	TotalPrice      float32
	TransactionUUID string
	PaymentMethod   PaymentMethod
	Status          OrderStatus
}

type OrderCreateParam struct {
	UserUUID      string
	PartUuids     []string
}

type OrderUpdateParam struct {
	OrderId         string
	TransactionUUID string
	PaymentMethod   PaymentMethod
	Status          OrderStatus
}
