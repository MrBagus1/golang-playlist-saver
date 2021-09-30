package web

import "time"

type PlaylistEditResponse struct {
	Name      string    `json:"name"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}
