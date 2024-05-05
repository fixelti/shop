package customError

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)
