package product

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	customError "shop/internal/common/errors"
	"shop/internal/common/models"
	queries "shop/internal/database/postgres/product/internal"
	"shop/internal/lib/database/postgres"
)

func (product Product) Create(ctx context.Context, productData models.ProductEntity) (uint, error) {
	tx, err := product.database.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, fmt.Errorf(customError.ErrBeginTransaction.Error(), err)
	}

	res, err := tx.Query(
		ctx,
		queries.CREATE,
		productData.Name,
		productData.Description,
		productData.Price,
		productData.ImageURL,
	)
	defer res.Close()
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return 0, fmt.Errorf(customError.ErrRollbackTransaction.Error(), rollbackErr)
		}
		return 0, fmt.Errorf(customError.ErrQuery.Error(), err)
	}

	createdProduct, err := postgres.ScanInStruct[models.ProductEntity](res)
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return 0, fmt.Errorf(customError.ErrRollbackTransaction.Error(), rollbackErr)
		}
		return 0, fmt.Errorf(customError.ErrScanInStruct.Error(), err)
	}

	if createdProduct == nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return 0, fmt.Errorf(customError.ErrRollbackTransaction.Error(), rollbackErr)
		}
		return 0, customError.ErrProductIsEmpty
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf(customError.ErrCommitTransaction.Error(), err)
	}
	return createdProduct.ID, nil
}
