package models

import "time"

type User struct {
	Id          int
	Name        string
	Identity    []byte
	Certificate []byte
	OrderSecret []byte
	CreatedAt   time.Time
}
