package user

import (
	"context"
	"shop/internal/lib/logger"
)

type userService interface {
	SignUp(ctx context.Context, email, password string) (uint, error)
}
type User struct {
	logger  logger.Logger
	service userService
}

func New(logger logger.Logger, service userService) User {
	return User{
		service: service,
		logger:  logger,
	}
}
