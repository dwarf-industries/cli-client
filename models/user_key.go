package models

import (
	"time"
)

type UserKey struct {
	Id        int
	UserId    int
	KeyData   string
	CreatedAt time.Time
}
