package user

import (
	"context"
	"github.com/golang-jwt/jwt"
	"shop/internal/common/models"
	"shop/internal/config"
	"shop/internal/lib/logger"
)

type userService interface {
	SignUp(ctx context.Context, email, password string) (uint, error)
	Login(ctx context.Context, email, password string) (models.AuthorizationTokens, error)
	RefreshAccessToken(ctx context.Context, claims jwt.MapClaims) (string, error)
}
type User struct {
	config  config.Config
	logger  logger.Logger
	service userService
}

func New(config config.Config, logger logger.Logger, service userService) User {
	return User{
		config:  config,
		service: service,
		logger:  logger,
	}
}
