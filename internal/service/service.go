package service

import (
	"github.com/plutonio00/pay-api/internal/repository"
)

type Users interface {
	Replenish(userId int, sum int) error
	Transfer(senderId int, recipientId int, sum int) error
}

type Services struct {
	Users Users
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	return &Services{
		Users: NewUsersService(deps.Repos.Users),
	}
}
