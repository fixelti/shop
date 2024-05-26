package user

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	customError "shop/internal/common/errors"
	"shop/internal/common/models"
	"shop/internal/lib/logger"
)

func (user User) SignUp(ctx context.Context, email, password string) (uint, error) {
	ctx = logger.WithOP(ctx, "service.user.SignUp")
	foundUser, err := user.db.GetByEmail(ctx, email)
	if err != nil {
		user.logger.Error(ctx, errors.Wrap(err, customError.ErrGetUserByEmail.Error()))
		return 0, err
	}

	if exist(foundUser) {
		return 0, customError.ErrUserExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		user.logger.Error(ctx, errors.Wrap(err, customError.ErrGenerateFromPassword.Error()))
		return 0, err
	}

	userID, err := user.db.Create(ctx, email, string(passwordHash), models.USER_ROLE)
	if err != nil {
		user.logger.Error(ctx, errors.Wrap(err, customError.ErrCreateUser.Error()))
		return 0, err
	}

	return userID, nil
}
