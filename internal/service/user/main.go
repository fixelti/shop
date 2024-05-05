package user

import (
	"context"
	"github.com/golang-jwt/jwt"
	"shop/internal/common/models"
	"shop/internal/lib/logger"
)

type userRepository interface {
	Create(ctx context.Context, email, password string) (uint, error)
	GetByEmail(ctx context.Context, email string) (models.UserEntity, error)
}

type generateJWTToken interface {
	GenerateTokens(ctx context.Context, userID uint) (models.AuthorizationTokens, error)
	RefreshToken(ctx context.Context, claims jwt.MapClaims) (string, error)
}

type User struct {
	logger logger.Logger
	db     userRepository
	jwt    generateJWTToken
}

func New(logger logger.Logger, db userRepository, jwt generateJWTToken) User {
	return User{
		logger: logger,
		db:     db,
		jwt:    jwt,
	}
}
