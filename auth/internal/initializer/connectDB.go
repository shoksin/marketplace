package initializer

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var DB *sqlx.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("error loading .env file")
	}

	var err error
	DB, err = sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatalf("error sqlx.Connect: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("error ping database: %v", err)
	}

	log.Println("Database connected.")
}
