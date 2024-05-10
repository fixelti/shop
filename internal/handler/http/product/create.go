package product

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shop/internal/common/models"
)

type CreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       uint   `json:"price" validate:"required"`
	ImageURL    string `json:"image_url" validate:"required"`
}

type CreateResponse struct {
	ID uint `json:"id"`
}

func (product Product) Create(c echo.Context) error {
	request := new(CreateRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	productID, err := product.service.Create(c.Request().Context(), models.ProductEntity{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		ImageURL:    request.ImageURL,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, CreateResponse{ID: productID})
}
