package repositories

import (
	"context"
	"database/sql"

	"github.com/tejashwinn/splitwise/types"
)

// UserRepositoryPG implements UserRepository using PostgreSQL
type UserRepositoryPG struct {
	DB *sql.DB
}

// NewUserRepository initializes a PostgreSQL repository
func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryPG{DB: db}
}

// GetAllUsers retrieves all users from the database
func (repo *UserRepositoryPG) GetAllUsers(ctx context.Context) ([]types.User, error) {
	rows, err := repo.DB.QueryContext(ctx, "SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []types.User
	for rows.Next() {
		var user types.User
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
