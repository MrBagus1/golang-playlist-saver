package web

import (
	"time"
)

type UserRegisterResponse struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
	Role     string    `json:"role"`
	Status   Status    `json:"status"`
}

type Status struct {
	Name      string    `json:"status_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
