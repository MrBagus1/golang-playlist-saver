package record

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

type Status struct {
	Id        int    `gorm:"primaryKey,not null, autoIncrement"`
	Name      string `gorm:"type:enum('FREE','PREMIUM');default:'FREE'"`
	UserId    int    `gorm:"not null"`
	TokenId   string
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiredAt null.Time
}
