package product

import (
	"context"
	"github.com/pkg/errors"
	customError "shop/internal/common/errors"
	"shop/internal/common/models"
	"shop/internal/lib/logger"
)

func (product Product) GetByID(ctx context.Context, id uint) (models.ProductEntity, error) {
	ctx = logger.WithOP(ctx, "service.product.GetByID")

	foundProduct, err := product.db.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, customError.ErrProductNotFound) {
			return models.ProductEntity{}, err
		}
		product.logger.Error(ctx, errors.Wrap(err, customError.ErrGetByIDProduct.Error()))
		return models.ProductEntity{}, err
	}

	return foundProduct, nil
}
