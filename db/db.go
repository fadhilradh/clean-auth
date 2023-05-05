package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("postgres", "postgresql://asfjtcdi:m5WzMP3nfg8Gi_4pxoAyFjgNCm8F2Vo5@babar.db.elephantsql.com:5432/asfjtcdi?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &Database{
		db: db,
	}, nil
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
