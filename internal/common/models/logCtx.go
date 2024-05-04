package models

import "github.com/jackc/pgx/v5"

type LogCtx struct {
	UserID           uint
	OP               string   // место вызова логов
	PostgresQueryRes pgx.Rows // возвращаемые данные из query запроса
}
