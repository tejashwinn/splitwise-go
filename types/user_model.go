package types

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int64        `json:"id"`
	Name      string       `json:"name"`
	Username  string       `json:"username"`
	Email     string       `json:"email"`
	Password  string       `json:"-"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt sql.NullTime `json:"updatedAt"`
}

type UserReq struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginReq struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type TokenRes struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserRes struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}
