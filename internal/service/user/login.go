package user

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	customError "shop/internal/common/errors"
	"shop/internal/common/models"
	"shop/internal/lib/logger"
)

const op = "service.user.Login"

func (user User) Login(ctx context.Context, email, password string) (models.AuthorizationTokens, error) {
	ctx = logger.WithOP(ctx, op)
	foundUser, err := user.db.GetByEmail(ctx, email)
	if err != nil {
		user.logger.Error(ctx, errors.Wrap(err, customError.ErrGetUserByEmail.Error()))
		return models.AuthorizationTokens{}, err
	}

	if !exist(foundUser) {
		return models.AuthorizationTokens{}, customError.ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password)); err != nil {
		return models.AuthorizationTokens{}, customError.ErrWrongPassword
	}

	tokens, err := user.jwt.GenerateTokens(ctx, foundUser.ID, foundUser.Role)
	if err != nil {
		user.logger.Error(ctx, errors.Wrap(err, customError.ErrGenerateTokens.Error()))
	}

	return tokens, nil
}
