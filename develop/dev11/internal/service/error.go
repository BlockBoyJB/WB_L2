package service

import "errors"

var (
	ErrUserNotFound = errors.New("user not exist")

	ErrEventNotFound     = errors.New("event not found")
	ErrEventAlreadyExist = errors.New("event already exist")
)
