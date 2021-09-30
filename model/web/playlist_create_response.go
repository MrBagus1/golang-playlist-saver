package web

import (
	"gorm.io/gorm"
	"time"
)

type PlaylistCreateResponse struct {
	Name      string    `json:"name"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt`json:"deleted_at"`
}
