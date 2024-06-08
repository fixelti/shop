package product

import (
	"context"
	"fmt"
	customError "shop/internal/common/errors"
	"shop/internal/common/models"
	queries "shop/internal/database/postgres/product/internal"

	"github.com/jackc/pgx/v5"
)

func (product Product) Update(ctx context.Context, updatedPoduct models.ProductEntity) error {
	tx, err := product.database.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf(customError.ErrBeginTransaction.Error(), err)
	}

	_, err = tx.Exec(
		ctx,
		queries.Update,
		updatedPoduct.Name,
		updatedPoduct.Description,
		updatedPoduct.Price,
		updatedPoduct.ImageURL,
		updatedPoduct.ID,
	)

	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return fmt.Errorf(customError.ErrRollbackTransaction.Error(), rollbackErr)
		}
		return fmt.Errorf(customError.ErrQuery.Error(), err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf(customError.ErrCommitTransaction.Error(), err)
	}

	return nil
}
