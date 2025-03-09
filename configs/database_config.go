package configs

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/tejashwinn/splitwise/types"
)

func ConnectDB(cfg *types.Config) (*sql.DB, error) {

	db, err := sql.Open("postgres", cfg.Db.Url)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetConnMaxIdleTime(1 * time.Minute)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
