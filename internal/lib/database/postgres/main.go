package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type Database struct {
	*pgxpool.Pool
}

func New(ctx context.Context, dsn string) Database {
	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	if err := conn.Ping(ctx); err != nil {
		log.Fatalf("failed to ping database: %s", err)
	}
	return Database{conn}
}

// ScanInStruct функция помещает ответ базы данных в структуру
func ScanInStruct[T any](rows pgx.Rows) (*T, error) {
	return pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[T])
}
