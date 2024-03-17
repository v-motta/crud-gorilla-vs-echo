package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := "postgres://mottinha:123456@127.0.0.1/postgres?sslmode=disable"
	return sql.Open("postgres", connStr)
}
