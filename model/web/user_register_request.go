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
	Dob      time.Time `json:"dob"`
	Gender   string    `json:"gender"`
}

func (u UserRegisterRequest) Validate() error {
	return validation.ValidateStruct(&u,
		//validasi Name tidak boleh kosong
		validation.Field(&u.Name, validation.Required),
		//Validasi email regex, dan tidak boleh kosong
		validation.Field(&u.Email, validation.Required, is.Email),
		//Validasi password tidak boleh kosong, dan minimal panjang 6 dan max 12
		validation.Field(&u.Password, validation.Required, validation.Length(6, 12)),
	)
}
