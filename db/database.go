package db

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("pgx", "postgres://admin:admin@localhost:5432/go_lesson")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(1 * time.Minute)
	db.SetConnMaxLifetime(5 * time.Minute)
	return db
}
