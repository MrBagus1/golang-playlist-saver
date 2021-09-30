package servtoken

import (
	"time"
)

type Token struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//func (u *Token) ToRecordToken() record.Token {
//	recordToken := record.Token{}
//	recordToken.Id = u.Id
//	recordToken.CreatedAt = u.CreatedAt
//	recordToken.UpdatedAt = u.UpdatedAt
//	return recordToken
//}