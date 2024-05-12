package customError

import "errors"

var (
	ErrBeginTransaction    = errors.New("failed to begin transaction")
	ErrRollbackTransaction = errors.New("failed to rollback transaction")
	ErrCommitTransaction   = errors.New("failed to commit transaction")
	ErrQuery               = errors.New("failed to query request")
	ErrScanInStruct        = errors.New("failed to scan in struct")
	ErrUserIsEmpty         = errors.New("user is empty")
	ErrProductIsEmpty      = errors.New("product is empty")
	ErrProductNotFound     = errors.New("product not found")
)
