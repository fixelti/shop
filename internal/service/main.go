package service

import (
	"shop/internal/database/postgres"
	"shop/internal/lib/logger"
	"shop/internal/service/user"
)

type Manager struct {
	logger          logger.Logger
	postgresManager postgres.Manager
	User            user.User
}

func New(logger logger.Logger, postgresManager postgres.Manager) Manager {
	return Manager{
		logger:          logger,
		postgresManager: postgresManager,
		User:            user.New(logger, postgresManager.User),
	}
}
