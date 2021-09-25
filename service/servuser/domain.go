package servuser

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gopkg.in/guregu/null.v4"
	"playlist-saver/model/record"
	"time"
)

type User struct {
	Id        int
	Name      string
	Password  string
	Email     string
	Birthday  time.Time
	Gender    string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    Status
}

type Status struct {
	Id        int
	Name      string
	UserId    int
	TokenId   string
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiredAt null.Time
}

func (u *User) ToRecordUser() record.User {
	recordUser := record.User{}
	recordUser.Name = u.Name
	recordUser.Email = u.Email
	recordUser.Birthday = u.Birthday
	recordUser.Role = u.Role
	recordUser.Status.Name = u.Status.Name
	recordUser.Status.CreatedAt = u.Status.CreatedAt
	recordUser.Status.UpdatedAt = u.Status.UpdatedAt
	return recordUser
}

func (u *User) FromRecordUser(recordUser record.User) {
	u.Name = recordUser.Name
	u.Email = recordUser.Email
	u.Birthday = recordUser.Birthday
	u.Role = recordUser.Role
	u.Status.Name = recordUser.Status.Name
	u.Status.CreatedAt = recordUser.Status.CreatedAt
	u.Status.UpdatedAt = recordUser.Status.UpdatedAt
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		//validasi Name tidak boleh kosong
		validation.Field(&u.Name, validation.Required),
		//Validasi email regex, dan tidak boleh kosong
		validation.Field(&u.Email, validation.Required, is.Email),
		//Validasi password tidak boleh kosong, dan minimal panjang 6 dan max 12
		validation.Field(&u.Password, validation.Required, validation.Length(6, 12)),
	)
}
