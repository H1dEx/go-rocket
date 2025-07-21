package model

import "errors"

var ErrOrderIdInvalid = errors.New("field order_uuid is invalid") 
var ErrUserIdInvalid = errors.New("field user_uuid is invalid") 
var ErrPaymentMethodInvalid = errors.New("field payment_method is invalid") 
