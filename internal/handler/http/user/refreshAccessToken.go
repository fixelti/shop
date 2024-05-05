package user

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"shop/internal/common/models"
)

type RefreshAccessTokenResponse models.AuthorizationTokens

func (user User) RefreshAccessToken(c echo.Context) error {
	claims, ok := c.Get("claims").(*jwt.MapClaims)
	if !ok {
		//TODO: возможно сделать вывод claims в логах
		return c.JSON(http.StatusUnauthorized, nil)
	}

	accessToken, err := user.service.RefreshAccessToken(c.Request().Context(), *claims)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, RefreshAccessTokenResponse{AccessToken: accessToken})
}
