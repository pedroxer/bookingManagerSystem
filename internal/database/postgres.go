package database

import (
	"database/sql"
	"fmt"
	"github.com/pedroxer/BookingManagerSystem/internal/configs"
)

func ConnToPostgres(pgconf *configs.Postgres) (*sql.DB, error) {
	DSN := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%d sslmode=%s",
		pgconf.Database,
		pgconf.Username,
		pgconf.Password,
		pgconf.Host,
		pgconf.Port,
		pgconf.Sslmode,
	)

	db, _ := sql.Open("postgres", DSN)
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ConnToPostgres: %w", err)
	}
	return db, nil
}
