package record

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

type Status struct {
	Id        int       `json:"id" gorm:"primaryKey,not null, autoIncrement"`
	Name      string    `json:"name" gorm:"type:enum('FREE','PREMIUM');default:'USER'"`
	UserId    int       `json:"user_id" gorm:"not null"`
	TokenId   string    `json:"token_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ExpiredAt null.Time `json:"expired_at"`
}
