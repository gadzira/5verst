package database

import (
	"github.com/5verst/config"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func MustConnectDB(cfg config.Config) *sqlx.DB {
	db, err := sqlx.Connect("mysql", cfg.DBDsn)
	if err != nil {
		log.Panic().Err(err).Msg("Can not connect to DB")
	}

	return db
}
