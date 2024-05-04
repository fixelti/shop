DB_USER ?= test
DB_PASSWORD ?= test
DB_NAME ?= shop-test
DB_PORT ?= 5555
DB_URL ?= localhost

# Миграции для базы данных
migration_up:
	goose -dir internal/database/postgres/migrations postgres "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_URL):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

migration_down:
	goose -dir internal/database/postgres/migrations postgres "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_URL):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down