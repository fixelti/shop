package product

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetListRequest struct {
	PageNum  uint `json:"page_num"`
	PageSize uint `json:"page_size"`
}

func (product Product) GetList(c echo.Context) error {
	request := new(GetListRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	foundProducts, err := product.service.GetList(c.Request().Context(), request.PageNum, request.PageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, foundProducts)
}
