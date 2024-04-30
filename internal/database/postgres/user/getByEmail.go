package user

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	customError "shop/internal/common/errors"
	"shop/internal/common/models"
	queries "shop/internal/database/postgres/user/internal"
	"shop/internal/lib/database/postgres"
)

func (user User) GetByEmail(ctx context.Context, email string) (models.UserEntity, error) {
	tx, err := user.database.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return models.UserEntity{}, customError.ErrBeginTransaction
	}

	res, err := tx.Query(
		ctx,
		queries.GetByEmail,
		email,
	)
	defer res.Close()
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return models.UserEntity{}, customError.ErrRollbackTransaction
		}
		return models.UserEntity{}, customError.ErrQuery
	}

	foundUser, err := postgres.ScanInStruct[models.UserEntity](res)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.UserEntity{}, nil
		}
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return models.UserEntity{}, customError.ErrRollbackTransaction
		}
		return models.UserEntity{}, customError.ErrScanInStruct
	}

	if foundUser == nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return models.UserEntity{}, customError.ErrRollbackTransaction
		}
		return models.UserEntity{}, customError.ErrUserIsEmpty
	}

	if err := tx.Commit(ctx); err != nil {
		return models.UserEntity{}, customError.ErrUserIsEmpty
	}

	return *foundUser, nil
}
