package user

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	customError "shop/internal/common/errors"
	"shop/internal/lib/logger"
)

func (user User) RefreshAccessToken(ctx context.Context, claims jwt.MapClaims) (string, error) {
	const op = "service.user.RefreshAccessToken"
	ctx = logger.WithOP(ctx, op)

	accessToken, err := user.jwt.RefreshToken(ctx, claims)
	if err != nil {
		user.logger.Error(ctx, errors.Wrap(err, customError.ErrGenerateRefreshToken.Error()))
	}
	return accessToken, nil
}
