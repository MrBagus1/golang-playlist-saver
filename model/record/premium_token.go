package record

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

type Token struct {
	Id        null.Int `gorm:"primaryKey, null , autoIncrement"`
	TokensId  string   `gorm:"size:255" json:"tokens_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
