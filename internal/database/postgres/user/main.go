package user

import (
	"shop/internal/lib/database/postgres"
)

type User struct {
	database postgres.Database
}

func New(database postgres.Database) *User {
	return &User{database: database}
}
