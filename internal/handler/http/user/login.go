package user

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	customError "shop/internal/common/errors"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func (user User) Login(c echo.Context) error {
	request := new(LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	tokens, err := user.service.Login(c.Request().Context(), request.Email, request.Password)
	if err != nil {
		if errors.Is(err, customError.ErrUserNotFound) {
			return c.JSON(http.StatusUnauthorized, nil)
		}
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, tokens)
}
