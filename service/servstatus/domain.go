package servstatus

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

type Status struct {
	Id        int
	Name      string
	UserId    int
	TokenId   null.String
	Token     Token
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiredAt int
}

type Token struct {
	Id        null.Int
	TokensId  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
