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

func (product Product) GetList(ctx context.Context, pageNum, pageSize uint) ([]models.ProductEntity, error) {
	tx, err := product.database.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return []models.ProductEntity{}, fmt.Errorf(customError.ErrBeginTransaction.Error(), err)
	}

	offset := (pageNum - 1) * pageSize
	res, err := tx.Query(
		ctx,
		queries.GetList,
		pageSize,
		offset,
	)
	defer res.Close()
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return []models.ProductEntity{}, fmt.Errorf(customError.ErrRollbackTransaction.Error(), rollbackErr)
		}
		return []models.ProductEntity{}, fmt.Errorf(customError.ErrQuery.Error(), err)
	}

	foundProducts, err := postgres.ScanInArrayStruct[models.ProductEntity](res)
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return []models.ProductEntity{}, fmt.Errorf(customError.ErrRollbackTransaction.Error(), rollbackErr)
		}
		return []models.ProductEntity{}, fmt.Errorf(customError.ErrScanInStruct.Error(), err)
	}

	return foundProducts, nil
}
