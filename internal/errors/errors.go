package errors

import "github.com/go-faster/errors"

var (
	UserNotFoundErr  = errors.New("user not found")
	WrongPasswordErr = errors.New("wrong password")
)
