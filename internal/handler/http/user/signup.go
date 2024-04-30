package user

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	customError "shop/internal/common/errors"
)

type SignUpRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type SignUpResponse struct {
	UserID uint `json:"user_id"`
}

func (user User) Signup(c echo.Context) error {
	request := new(SignUpRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	userID, err := user.service.SignUp(c.Request().Context(), request.Email, request.Password)
	if err != nil {
		if errors.Is(err, customError.ErrUserExists) {
			return c.JSON(http.StatusConflict, nil)
		}
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, SignUpResponse{UserID: userID})
}
