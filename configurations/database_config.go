package config

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/tejashwinn/splitwise/types"
)

func ConnectDB(cfg *types.Config) (*sql.DB, error) {

	db, err := sql.Open("postgres", cfg.Db.Url)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
