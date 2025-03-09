package types

import (
	"time"
)

type GroupUserMap struct {
	Id        int64     `json:"id"`
	GroupId   int64     `json:"groupId"`
	UserId    int64     `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy int64     `json:"createdBy"`
}
