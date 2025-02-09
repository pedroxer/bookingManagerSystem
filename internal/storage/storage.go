package storage

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

type Storage struct {
	pgDB   *sql.DB
	logger *logrus.Logger
}

func New(pgDB *sql.DB, logger *logrus.Logger) *Storage {
	return &Storage{
		pgDB:   pgDB,
		logger: logger,
	}
}
