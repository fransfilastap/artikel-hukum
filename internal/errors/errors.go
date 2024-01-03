package errors

import "errors"

var (
	ErrUserDoesNotExists    = errors.New("user doesn't exist")
	ErrEmailAlreadyExists   = errors.New("the email is already in use")
	ErrOldPasswordIncorrect = errors.New("incorrect old password")
)
