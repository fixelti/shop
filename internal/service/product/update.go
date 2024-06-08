package product

import (
	"context"
	customError "shop/internal/common/errors"
	"shop/internal/common/models"
	"shop/internal/lib/logger"

	"github.com/pkg/errors"
)

func (product Product) Update(ctx context.Context, updatedProduct models.ProductEntity) error {
	ctx = logger.WithOP(ctx, "service.product.Update")

	if err := product.db.Update(ctx, updatedProduct); err != nil {
		product.logger.Error(ctx, errors.Wrap(err, customError.ErrUpdateProduct.Error()))
		return err
	}

	return nil
}
