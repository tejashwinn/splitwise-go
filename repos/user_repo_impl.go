package repos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/tejashwinn/splitwise/mappers"
	"github.com/tejashwinn/splitwise/types"
	"github.com/tejashwinn/splitwise/utils"
)

type UserRepoImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepo {
	return &UserRepoImpl{DB: db}
}

func (repo *UserRepoImpl) FindAll(
	ctx context.Context,
) ([]types.User, error) {
	query := `
		SELECT OBJECT_ID,
			USR_NAME,
			USR_USERNAME,
			USR_PASSWORD,
			USR_EMAIL,
			CREATED_AT,
			UPDATED_AT
		FROM SW_USR
	`
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []types.User{}
	for rows.Next() {
		user, err := mappers.MapRowsToUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}
	return users, nil
}

func (repo *UserRepoImpl) Save(
	ctx context.Context,
	user *types.User,
) (*types.User, error) {
	query := `
		INSERT INTO SW_USR (
			OBJECT_ID,
			USR_NAME,
			USR_USERNAME,
			USR_EMAIL,
			USR_PASSWORD,
			CREATED_AT
		)
		VALUES (DEFAULT, $1, $2, $3, $4, $5)
		RETURNING OBJECT_ID
	`
	user.CreatedAt = time.Now()
	row := repo.DB.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Username,
		user.Email,
		user.Password,
		user.CreatedAt,
	)

	if row.Err() != nil {
		log.Println(row.Err())
		return user, errors.New("error during insertion")
	}
	row.Scan(&user.Id)
	return user, nil
}

func (repo *UserRepoImpl) FindByEmailOrUsername(
	ctx context.Context,
	usernameEmail string,
) (*types.User, error) {
	query := `
		SELECT OBJECT_ID,
			USR_NAME,
			USR_USERNAME,
			USR_PASSWORD,
			USR_EMAIL,
			CREATED_AT,
			UPDATED_AT
		FROM SW_USR
		WHERE USR_EMAIL = $1
			OR 
			USR_USERNAME = $1
	`

	row := repo.DB.QueryRowContext(
		ctx,
		query,
		usernameEmail,
	)

	if row.Err() != nil {
		log.Println(row.Err())
		return nil, errors.New("Unable to find user")
	}
	user, err := mappers.MapRowToUser(row)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to map user")
	}
	return user, nil
}

func (repo *UserRepoImpl) FindById(
	ctx context.Context,
	id int64,
) (*types.User, error) {
	query := `
		SELECT OBJECT_ID,
			USR_NAME,
			USR_USERNAME,
			USR_PASSWORD,
			USR_EMAIL,
			CREATED_AT,
			UPDATED_AT
		FROM SW_USR
		WHERE OBJECT_ID = $1
	`

	row := repo.DB.QueryRowContext(
		ctx,
		query,
		id,
	)

	if row.Err() != nil {
		log.Println(row.Err())
		return nil, errors.New("Unable to find user")
	}
	user, err := mappers.MapRowToUser(row)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to map user")
	}
	return user, nil
}

func (repo *UserRepoImpl) FindByIdIn(
	ctx context.Context,
	userIds []int64,
) ([]types.User, error) {

	placeholders, args := utils.GenerateSQLPlaceholders(len(userIds))

	for i, id := range userIds {
		args[i] = id
	}

	query := fmt.Sprintf(`
		SELECT OBJECT_ID,
			USR_NAME,
			USR_USERNAME,
			USR_PASSWORD,
			USR_EMAIL,
			CREATED_AT,
			UPDATED_AT
		FROM SW_USR
			WHERE OBJECT_ID IN (%s)
	`, placeholders)

	rows, err := repo.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []types.User{}
	for rows.Next() {
		user, err := mappers.MapRowsToUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}
	return users, nil
}
