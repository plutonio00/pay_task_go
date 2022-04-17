package model

import ()

type User struct {
	Id        uint  `json:"id" example:"1"`
	FirstName string `json:"first_name" example:"George"`
	LastName  string `json:"last_name" example:"Smith"`
	Balance   uint   `json:"balance" example:"10"`
}
