package error

import (
	"errors"
)

var (
	ErrUserNotFound   = errors.New("User doesn't exist")
	ErrNotEnoughMoney = errors.New("Not enough money on account")
)
