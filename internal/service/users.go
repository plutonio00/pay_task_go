package service

import (
	// 	api_errors "github.com/plutonio00/pay-api/internal/error"
	"github.com/plutonio00/pay-api/internal/repository"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {

	return &UsersService{
		repo: repo,
	}
}

func (s *UsersService) Replenish(userId int, sum int) error {
	return s.repo.Replenish(userId, sum)
}

func (s *UsersService) Transfer(senderId int, recipientId int, sum int) error {
	return s.repo.Transfer(senderId, recipientId, sum)
}
