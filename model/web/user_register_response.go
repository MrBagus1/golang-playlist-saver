package web

import (
	"playlist-saver/service/servuser"
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

func (response *UserRegisterResponse) FromDomainService(u servuser.User) {
	response.Name = u.Name
	response.Email = u.Email
	response.Birthday = u.Birthday
	response.Role = u.Role
	response.Status.Name = u.Status.Name
	response.Status.CreatedAt = u.Status.CreatedAt
	response.Status.UpdatedAt = u.Status.UpdatedAt
}
