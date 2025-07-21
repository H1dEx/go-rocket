package model

import "errors"

var ErrOrderNotFound = errors.New("order not found")
var ErrOrderAlreadyExists = errors.New("order with this id already exists")

var ErrPartsNotFound = errors.New("parts not found")
var ErrPartsWrongAmoundFound = errors.New("wrong amound of parts found")
