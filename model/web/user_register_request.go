package web

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"time"
)

type UserRegisterRequest struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Birthday time.Time `json:"birthday"`
	Gender   string    `json:"gender"`
}

func (request *UserRegisterRequest) ToDomainService() servuser.User {
	u := servuser.User{}

	u.Name = request.Name
	u.Password = request.Password
	u.Email = request.Email
	u.Birthday = request.Birthday
	u.Gender = request.Gender

	return u
}
