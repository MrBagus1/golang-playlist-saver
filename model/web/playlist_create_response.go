package web

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

type PlaylistCreateResponse struct {
	Name      string    `json:"name"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}
