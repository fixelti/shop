package product

import (
	"context"
	customError "shop/internal/common/errors"
	"shop/internal/common/models"
	"shop/internal/lib/logger"

	"github.com/pkg/errors"
)

func (product Product) GetList(ctx context.Context, pageNum uint, pageSize uint) ([]models.ProductEntity, error) {
	ctx = logger.WithOP(ctx, "service.product.GetList")

	foundProducts, err := product.db.GetList(ctx, pageNum, pageSize)
	if err != nil {
		product.logger.Error(ctx, errors.Wrap(err, customError.ErrGetByListProduct.Error()))
		return []models.ProductEntity{}, err
	}

	return foundProducts, nil
}
