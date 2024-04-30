package user

import (
	"shop/internal/lib/logger"
	"shop/internal/service/user"
)

type User struct {
	logger  logger.Logger
	service user.User
}

func New(logger logger.Logger, service user.User) User {
	return User{
		service: service,
		logger:  logger,
	}
}
