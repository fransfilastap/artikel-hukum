package v1

import "errors"

var (
	ErrUserDoesNotExists  = errors.New("user doesn't exist")
	ErrEmailAlreadyExists = errors.New("the email is already in use")
)
