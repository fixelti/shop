package user

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	customError "shop/internal/common/errors"
	"shop/internal/common/models"
	queries "shop/internal/database/postgres/user/internal"
	"shop/internal/lib/database/postgres"
)

func (user User) Create(ctx context.Context, email, password string, role models.Role) (uint, error) {
	tx, err := user.database.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, fmt.Errorf(customError.ErrBeginTransaction.Error(), err)
	}

	res, err := tx.Query(
		ctx,
		queries.CREATE,
		email,
		password,
		role,
	)
	defer res.Close()
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return 0, fmt.Errorf(customError.ErrRollbackTransaction.Error(), rollbackErr)
		}
		return 0, fmt.Errorf(customError.ErrQuery.Error(), err)
	}

	createdUser, err := postgres.ScanInStruct[models.UserEntity](res)
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return 0, fmt.Errorf(customError.ErrRollbackTransaction.Error(), rollbackErr)
		}
		return 0, fmt.Errorf(customError.ErrScanInStruct.Error(), err)
	}

	if createdUser == nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return 0, fmt.Errorf(customError.ErrRollbackTransaction.Error(), rollbackErr)
		}
		return 0, customError.ErrUserIsEmpty
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf(customError.ErrCommitTransaction.Error(), err)
	}
	return createdUser.ID, nil
}
