package main

import (
	"context"
	"fmt"
	"log/slog"
	"shop/internal/config"
	postgresDB "shop/internal/database/postgres"
	"shop/internal/handler/http"
	"shop/internal/lib/database/postgres"
	"shop/internal/lib/jwt"
	"shop/internal/lib/logger"
	"shop/internal/service"
)

const (
	LocalEnv = "local"
	DevEnv   = "dev"
	ProdEnv  = "prod"
)

func main() {
	cfg := config.MustGetConfig("./config")
	log := logger.New(mustGetEnvironment(string(cfg.Env)))
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Name)
	db := postgres.New(context.Background(), dsn)

	jwtGenerator := jwt.New(
		cfg.JWT.AccessTokenLifeTime,
		cfg.JWT.RefreshTokenLifeTime,
		cfg.JWT.AccessTokenKey,
		cfg.JWT.RefreshTokenKey)

	postgresManager := postgresDB.New(db)
	serviceManager := service.New(jwtGenerator, log, postgresManager)

	server := http.New(cfg, log, serviceManager)

	log.Info(context.Background(), fmt.Sprintf("starting server on %s port", cfg.Server.Port))
	if err := server.Start(cfg.Server.Port); err != nil {
		panic(err)
	}
}

func mustGetEnvironment(env string) slog.Level {
	switch env {
	case LocalEnv:
		return slog.LevelInfo
	case DevEnv:
		return slog.LevelWarn
	case ProdEnv:
		return slog.LevelError
	default:
		panic("unknown environment")
	}
}
