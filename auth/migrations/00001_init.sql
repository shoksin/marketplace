-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    user_id    UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username   VARCHAR(100)   NOT NULL,
    email      VARCHAR(150)   NOT NULL,
    password   VARCHAR(200)   NOT NULL,
    birthday   VARCHAR(10)    NOT NULL,
    created_at TIMESTAMP      NOT NULL,
    updated_at TIMESTAMP      NULL,
    deleted_at TIMESTAMP      NULL
);

CREATE TABLE IF NOT EXISTS admins (
    admin_id   UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username   VARCHAR(100)   NOT NULL,
    password   VARCHAR(200)   NOT NULL,
    created_at TIMESTAMP      NOT NULL,
    updated_at TIMESTAMP      NULL,
    deleted_at TIMESTAMP      NULL
);

-- +goose Down
DROP TABLE IF EXISTS admins;
DROP TABLE IF EXISTS users;
DROP EXTENSION IF EXISTS "uuid-ossp";
