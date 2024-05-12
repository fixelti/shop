package product

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	customError "shop/internal/common/errors"
	"shop/internal/common/models"
	queries "shop/internal/database/postgres/product/internal"
	"shop/internal/lib/database/postgres"
)

func (product Product) GetByID(ctx context.Context, id uint) (models.ProductEntity, error) {
	tx, err := product.database.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return models.ProductEntity{}, fmt.Errorf(customError.ErrBeginTransaction.Error(), err)
	}

	res, err := tx.Query(
		ctx,
		queries.GetByID,
		id,
	)
	defer res.Close()
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return models.ProductEntity{}, fmt.Errorf(customError.ErrRollbackTransaction.Error(), rollbackErr)
		}
		return models.ProductEntity{}, fmt.Errorf(customError.ErrQuery.Error(), err)
	}

	foundProduct, err := postgres.ScanInStruct[models.ProductEntity](res)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.ProductEntity{}, customError.ErrProductNotFound
		}
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return models.ProductEntity{}, fmt.Errorf(customError.ErrRollbackTransaction.Error(), rollbackErr)
		}
		return models.ProductEntity{}, fmt.Errorf(customError.ErrScanInStruct.Error(), err)
	}

	if foundProduct == nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return models.ProductEntity{}, fmt.Errorf(customError.ErrRollbackTransaction.Error(), rollbackErr)
		}
		return models.ProductEntity{}, customError.ErrUserIsEmpty
	}

	if err := tx.Commit(ctx); err != nil {
		return models.ProductEntity{}, fmt.Errorf(customError.ErrCommitTransaction.Error(), err)
	}

	return *foundProduct, nil
}
