package jwt

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	customError "shop/internal/common/errors"
	"shop/internal/common/models"
	"time"
)

const (
	ClaimsUserIDKey     = "user_id"
	ClaimsExpirationKey = "expiration"
)

type JWT struct {
	expirationAccessToken  time.Duration
	expirationRefreshToken time.Duration
	accessTokenKey         string
	refreshTokenKey        string
}

func New(
	expirationAccessToken time.Duration,
	expirationRefreshToken time.Duration,
	accessTokenKey string,
	refreshTokenKey string,
) JWT {
	return JWT{
		expirationAccessToken:  expirationAccessToken,
		expirationRefreshToken: expirationRefreshToken,
		accessTokenKey:         accessTokenKey,
		refreshTokenKey:        refreshTokenKey,
	}
}

func (j JWT) GenerateTokens(ctx context.Context, userID uint) (models.AuthorizationTokens, error) {
	var err error
	token := models.AuthorizationTokens{}
	// claims для access токена
	claims := jwt.MapClaims{
		ClaimsUserIDKey:     userID,
		ClaimsExpirationKey: time.Now().Add(j.expirationAccessToken).Unix(),
	}

	token.AccessToken, err = j.generateToken(ctx, j.accessTokenKey, claims)
	if err != nil {
		return token, errors.Wrap(err, customError.ErrGenerateAccessToken.Error())
	}

	// claims для refresh токена
	claims[ClaimsExpirationKey] = time.Now().Add(j.expirationRefreshToken).Unix()
	token.RefreshToken, err = j.generateToken(ctx, j.refreshTokenKey, claims)
	if err != nil {
		return token, errors.Wrap(err, customError.ErrGenerateRefreshToken.Error())
	}

	return token, nil
}

func (j JWT) RefreshToken(ctx context.Context, claims jwt.MapClaims) (string, error) {
	var err error
	var token string

	claims[ClaimsExpirationKey] = time.Now().Add(j.expirationAccessToken).Unix()
	token, err = j.generateToken(ctx, j.accessTokenKey, claims)
	if err != nil {
		return token, errors.Wrap(err, customError.ErrGenerateAccessToken.Error())
	}

	return token, nil
}

func (j JWT) generateToken(ctx context.Context, tokenKey string, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token == nil {
		return "", customError.ErrTokenWithClaimsIsNil
	}

	return token.SignedString([]byte(tokenKey))
}
