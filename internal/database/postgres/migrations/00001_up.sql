-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(30) DEFAULT 'no_data',
    surname VARCHAR(30) DEFAULT 'no_data',
    patronymic VARCHAR(30) DEFAULT 'no_data',
    date_of_birth INTEGER DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products (
     id SERIAL PRIMARY KEY,
     name VARCHAR(100) NOT NULL,
     description VARCHAR(255) NOT NULL,
     price INTEGER DEFAULT 0,
     image_url VARCHAR(30) DEFAULT 'no_data',
     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
     deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS products;
-- +goose StatementEnd