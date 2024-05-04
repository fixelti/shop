package customError

import "errors"

var (
	ErrGetUserByEmail       = errors.New("failed to get user by email")
	ErrGenerateFromPassword = errors.New("failed to generate from password")
	ErrUserExists           = errors.New("user already exists")
	ErrCreateUser           = errors.New("failed to create user")
)
