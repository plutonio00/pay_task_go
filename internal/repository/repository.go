package repository

import (
	"database/sql"
)

type Users interface {
	Replenish(userId int, sum int) error
	Transfer(senderId int, recipientId int, sum int) error
}

type Repositories struct {
	Users Users
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Users: NewUsersRepo(db),
	}
}
