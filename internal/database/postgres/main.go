package postgres

import (
	"shop/internal/database/postgres/user"
	"shop/internal/lib/database/postgres"
)

type Manager struct {
	database postgres.Database
	User     user.User
}

func New(database postgres.Database) Manager {
	return Manager{
		database: database,
		User:     user.New(database),
	}
}
