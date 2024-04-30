package models

import "github.com/jackc/pgx/v5"

type LogCtx struct {
	UserID           uint
	OP               string
	PostgresQueryRes pgx.Rows
}
