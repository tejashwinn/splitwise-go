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

type GroupUserRepoImpl struct {
	DB      *sql.DB
	JwtUtil *utils.JwtUtil
}

func NewGroupUserRepository(
	db *sql.DB,
	jwtUtil *utils.JwtUtil,
) GroupUserRepo {
	return &GroupUserRepoImpl{DB: db, JwtUtil: jwtUtil}
}

func (repo *GroupUserRepoImpl) Save(
	ctx context.Context,
	groupUser *types.GroupUser,
) (*types.GroupUser, error) {
	var err error
	query := `
		INSERT INTO SW_GRP_USR (
			OBJECT_ID,
			GRP_ID,
			USR_ID,
			CREATED_AT,
			CREATED_BY_USR_ID
		)
		VALUES (DEFAULT, $1, $2, $3, $4)
		RETURNING OBJECT_ID
	`
	groupUser.CreatedAt = time.Now()
	groupUser.CreatedBy, err = repo.JwtUtil.GetUserId(ctx)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	row := repo.DB.QueryRowContext(
		ctx,
		query,
		groupUser.GroupId,
		groupUser.UserId,
		groupUser.CreatedAt,
		groupUser.CreatedBy,
	)

	if row.Err() != nil {
		log.Println(row.Err())
		return nil, errors.New("error during insertion")
	}
	row.Scan(&groupUser.Id)
	return groupUser, nil
}

func (repo *GroupUserRepoImpl) FindUserIdByGroupId(ctx context.Context, groupId int64) ([]int64, error) {
	var err error
	query := `
		SELECT USR_ID
		FROM SW_GRP_USR
		WHERE GRP_ID = $1
	`

	rows, err := repo.DB.QueryContext(
		ctx,
		query,
		groupId,
	)
	if err != nil {
		return nil, errors.New("Unable to fetch users")
	}
	defer rows.Close()
	users := []int64{}
	for rows.Next() {
		var user int64
		rows.Scan(&user)
		users = append(users, user)
	}
	return users, nil
}
