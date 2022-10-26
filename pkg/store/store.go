package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type DBStore struct {
	db *sqlx.DB
}

func NewDBStore(db *sqlx.DB) (*DBStore, error) {
	return &DBStore{
		db: db,
	}, nil
}

func MustDBStore(db *sqlx.DB) DBStore {
	dbs, err := NewDBStore(db)
	if err != nil {
		log.Panic().Msgf("Failed init config: %v", err)
	}

	return *dbs
}

func (s *DBStore) Ping() error {
	return s.db.Ping()
}
