package product

import (
	"context"
	"shop/internal/common/models"
	"shop/internal/lib/logger"
)

type productService interface {
	Create(ctx context.Context, productData models.ProductEntity) (uint, error)
	GetByID(ctx context.Context, id uint) (models.ProductEntity, error)
}
type Product struct {
	logger  logger.Logger
	service productService
}

func New(logger logger.Logger, service productService) Product {
	return Product{
		service: service,
		logger:  logger,
	}
}
