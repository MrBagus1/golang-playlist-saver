package web

import "time"

type UserUpdateResponse struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
	Gender   string    `json:"gender"`
}