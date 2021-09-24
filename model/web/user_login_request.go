package web

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u UserLoginRequest) Validate() error {
	return validation.ValidateStruct(&u,
		//validasi Name tidak boleh kosong
		//Validasi email regex, dan tidak boleh kosong
		validation.Field(&u.Email, validation.Required, is.Email),
		//Validasi password tidak boleh kosong, dan minimal panjang 6 dan max 12
		validation.Field(&u.Password, validation.Required, validation.Length(6, 12)),
	)
}
