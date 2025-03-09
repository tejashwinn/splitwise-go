package repos

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/tejashwinn/splitwise/types"
	"github.com/tejashwinn/splitwise/utils"
)

type GroupRepoImpl struct {
	DB      *sql.DB
	JwtUtil *utils.JwtUtil
}

func NewGroupRepository(
	db *sql.DB,
	jwtUtil *utils.JwtUtil,
) GroupRepo {
	return &GroupRepoImpl{DB: db, JwtUtil: jwtUtil}
}

func (repo *GroupRepoImpl) Save(
	ctx context.Context,
	group *types.Group,
) (*types.Group, error) {
	var err error
	query := `
		INSERT INTO SW_GRP (
			OBJECT_ID,
			GRP_NAME,
			GRP_DESCRITPION,
			CUR_ID,
			CREATED_AT,
			CREATED_BY_USR_ID
		)
		VALUES (DEFAULT, $1, $2, $3, $4, $5)
		RETURNING OBJECT_ID
	`
	group.CreatedAt = time.Now()
	group.CreatedBy, err = repo.JwtUtil.GetUserId(ctx)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	row := repo.DB.QueryRowContext(
		ctx,
		query,
		group.Name,
		group.Description,
		group.CurrencyId,
		group.CreatedAt,
		group.CreatedBy,
	)

	if row.Err() != nil {
		log.Println(row.Err())
		return nil, errors.New("error during insertion")
	}
	row.Scan(&group.Id)
	return group, nil
}
