package web

import "time"

type UserRegisterResponse struct {
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Dob    time.Time `json:"dob"`
	Role   string    `json:"role"`
	Status status    `json:"status"`
}

type status struct {
	Name      string    `json:"status_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
