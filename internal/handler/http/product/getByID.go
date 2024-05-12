package product

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	customError "shop/internal/common/errors"
)

type GetProductByIDRequest struct {
	ID uint `query:"id" binding:"required"`
}

func (product Product) GetByID(c echo.Context) error {
	request := new(GetProductByIDRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	foundProduct, err := product.service.GetByID(c.Request().Context(), request.ID)
	if err != nil {
		if errors.Is(err, customError.ErrProductNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": customError.ErrProductNotFound.Error()})
		}
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, foundProduct.ToDTO())
}
