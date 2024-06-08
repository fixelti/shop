package product

import (
	"net/http"
	"shop/internal/common/models"

	"github.com/labstack/echo/v4"
)

type UpdateProductRequest struct {
	ID          uint   `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       uint   `json:"price" validate:"required"`
	ImageURL    string `json:"image_url" validate:"required"`
}

func (product Product) Update(c echo.Context) error {
	request := new(UpdateProductRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err := product.service.Update(c.Request().Context(), models.ProductEntity{
		ID:          request.ID,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		ImageURL:    request.ImageURL,
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, nil)
}
