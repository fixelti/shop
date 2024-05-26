package http

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	customError "shop/internal/common/errors"
	"shop/internal/common/models"
	"strings"
)

func (h Handler) VerifyAccessToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := GetToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, nil)
		}
		accessTokenKey := h.config.JWT.AccessTokenKey
		accessToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(accessTokenKey), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		if !accessToken.Valid {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		return next(c)
	}
}

func (h Handler) VerifyRefreshToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := GetToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		refreshTokenKey := h.config.JWT.RefreshTokenKey
		claims := new(jwt.MapClaims)
		refreshToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(refreshTokenKey), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		if !refreshToken.Valid {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		c.Set("claims", claims)
		return next(c)
	}
}

func (h Handler) AdministratorCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := GetToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, nil)
		}
		accessTokenKey := h.config.JWT.AccessTokenKey
		accessToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(accessTokenKey), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		if !accessToken.Valid {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		claims, ok := accessToken.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, nil)
		}
		userRole := models.Role(claims["role"].(string))
		if userRole != models.ADMIN_ROLE {
			return c.JSON(http.StatusForbidden, nil)
		}
		return next(c)
	}
}

func GetToken(c echo.Context) (string, error) {
	requestHeader := c.Request().Header.Get("Authorization")
	arr := strings.Split(requestHeader, "Bearer ")
	if len(arr) != 2 {
		return "", customError.ErrInvalidCredentials
	}

	return arr[1], nil
}
