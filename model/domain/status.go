package domain

import (
	"time"
)

type Status struct {
	Id        int       `json:"id" gorm:"primaryKey,not null, autoIncrement"`
	Name      string    `json:"name" `
	UserId    int       `json:"user_id" gorm:"not null"`
	TokenId   string    `json:"token_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ExpiredAt time.Time `json:"expired_at"`
}