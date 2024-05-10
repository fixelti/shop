package product

import (
	"context"
	"github.com/pkg/errors"
	customError "shop/internal/common/errors"
	"shop/internal/common/models"
	"shop/internal/lib/logger"
)

func (product Product) Create(ctx context.Context, productData models.ProductEntity) (uint, error) {
	ctx = logger.WithOP(ctx, "service.product.Create")

	productID, err := product.db.Create(ctx, productData)
	if err != nil {
		product.logger.Error(ctx, errors.Wrap(err, customError.ErrCreateProduct.Error()))
		return 0, err
	}

	return productID, nil
}
