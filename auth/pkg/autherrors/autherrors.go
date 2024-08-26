package autherrors

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrInvalidPlayerStatus = errors.New("invalid player status")
	ErrWrongPassword       = errors.New("wrong password")
)
