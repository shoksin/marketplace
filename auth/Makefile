DB_URL = postgres://postgres:postgres@localhost:5433/auth-service

migrate-create:
	goose -s -dir ./migrations create $(name) sql

migrate-up:
	goose -dir ./migrations postgres "$(DB_URL)" up

migrate-down:
	goose -dir ./migrations postgres "$(DB_URL)" down
