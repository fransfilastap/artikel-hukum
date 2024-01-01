package v1

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("the email is already in use")
)
