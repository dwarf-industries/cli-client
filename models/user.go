package models

import "time"

type User struct {
	Id          int
	Certificate []byte
	Name        string
	CreatedAt   time.Time
}
