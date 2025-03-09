package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/tejashwinn/splitwise/types"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (repo *UserRepositoryImpl) GetAllUsers(
	ctx context.Context,
) ([]types.User, error) {
	query := `
		SELECT 
		OBJECT_ID,
		USR_NAME,
		USR_PASSWORD,
		USR_EMAIL,
		CREATED_AT
		FROM SW_USR
	`
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []types.User{}
	for rows.Next() {
		user, err := mapUserRows(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo *UserRepositoryImpl) InsertOneUser(
	ctx context.Context,
	user *types.User,
) (types.User, error) {
	query := `
		INSERT INTO SW_USR (
			OBJECT_ID,
			USR_NAME,
			USR_EMAIL,
			USR_PASSWORD,
			CREATED_AT
		)
		VALUES (DEFAULT, $1, $2, $3, $4)
		RETURNING OBJECT_ID
	`
	user.CreatedAt = time.Now()
	row := repo.DB.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
	)

	if row.Err() != nil {
		log.Println(row.Err())
		return *user, errors.New("error during insertion")
	}
	row.Scan(&user.Id)
	return *user, nil
}

func mapUserRows(rows *sql.Rows) (types.User, error) {
	var user types.User
	if err := rows.Scan(
		&user.Id,
		&user.Name,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
	); err != nil {
		return user, err
	}
	return user, nil
}
