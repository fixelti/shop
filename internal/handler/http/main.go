package http

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"shop/internal/config"
	"shop/internal/handler/http/product"
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
	config    config.Config
	user      user.User
	product   product.Product
}

func New(config config.Config, logger logger.Logger, serviceManager service.Manager) *echo.Echo {
	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}
	http := Handler{
		echo:    e,
		config:  config,
		user:    user.New(config, logger, serviceManager.User),
		product: product.New(logger, serviceManager.Product),
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
