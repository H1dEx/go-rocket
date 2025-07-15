package model

import "errors"

var ErrOrderNotFound = errors.New("order not found")
var ErrOrderAlreadyExists = errors.New("order with this id already exists")
