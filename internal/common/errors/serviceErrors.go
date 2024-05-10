package customError

import "errors"

var (
	ErrGetUserByEmail       = errors.New("failed to get user by email")
	ErrGenerateFromPassword = errors.New("failed to generate from password")
	ErrUserExists           = errors.New("user already exists")
	ErrCreateUser           = errors.New("failed to create user")
	ErrUserNotFound         = errors.New("user not found")
	ErrWrongPassword        = errors.New("wrong password")
	ErrGenerateTokens       = errors.New("failed to generate tokens")
	ErrCreateProduct        = errors.New("failed to product user")
)
