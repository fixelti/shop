package product

import "shop/internal/lib/database/postgres"

type Product struct {
	database postgres.Database
}

func New(database postgres.Database) Product {
	return Product{database: database}
}
