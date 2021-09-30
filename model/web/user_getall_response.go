package web

import "time"

type UserGetAllResponse struct {
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Birthday  time.Time     `json:"birthday"`
	Role      string        `json:"role"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Status    UserAllStatus `json:"status"`
}

type UserAllStatus struct {
	Id        int       `json:"id"`
	Name      string    `json:"status_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
