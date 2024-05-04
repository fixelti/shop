package user

import (
	"context"
	"shop/internal/common/models"
	"shop/internal/lib/logger"
)

type userRepository interface {
	Create(ctx context.Context, email, password string) (uint, error)
	GetByEmail(ctx context.Context, email string) (models.UserEntity, error)
}

type User struct {
	logger logger.Logger
	db     userRepository
}

func New(logger logger.Logger, db userRepository) User {
	return User{
		logger: logger,
		db:     db,
	}
}
