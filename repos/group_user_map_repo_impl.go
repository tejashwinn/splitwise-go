package repos

import (
	"database/sql"
)

type GroupUserMapRepoImpl struct {
	DB *sql.DB
}

func NewGroupUserMapRepository(db *sql.DB) GroupUserMapRepo {
	return &GroupUserMapRepoImpl{DB: db}
}
