package product

import (
	"context"
	"shop/internal/common/models"
	"shop/internal/lib/logger"
)

type productRepository interface {
	Create(ctx context.Context, product models.ProductEntity) (uint, error)
}

type Product struct {
	logger logger.Logger
	db     productRepository
}

func New(logger logger.Logger, db productRepository) Product {
	return Product{
		logger: logger,
		db:     db,
	}
}
