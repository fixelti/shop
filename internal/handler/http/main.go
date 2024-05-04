package http

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"shop/internal/handler/http/user"
	"shop/internal/lib/logger"
	"shop/internal/service"
)

type Validator struct {
	validator *validator.Validate
}

type Handler struct {
	echo      *echo.Echo
	validator *echo.Validator
	user      user.User
}

func New(logger logger.Logger, serviceManager service.Manager) *echo.Echo {
	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}
	http := Handler{
		echo: e,
		user: user.New(logger, serviceManager.User),
	}

	http.handlers()
	return http.echo
}

func (validator *Validator) Validate(data interface{}) error {
	if err := validator.validator.Struct(data); err != nil {
		return err
	}
	return nil
}
