package domain

import (
	"gorm.io/gorm"
	"playlist-saver/model/web"
	"time"
)

type User struct {
	gorm.Model
	Id        int       `json:"id" gorm:"primary_key,autoIncrement,not null"`
	Name      string    `json:"name" gorm:"not null"`
	Password  string    `json:"password" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique"`
	Birthday  time.Time `json:"birthday"`
	Gender    string    `json:"gender"`
	Role      string    `json:"role" gorm:"type:enum('ADMIN','USER');default:'USER'"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    Status    `json:"status" gorm:"foreignKey:UserId"`
}

func (u *User) UserInsertRequest(request web.UserRegisterRequest) {
	u.Name = request.Name
	u.Password = request.Password
	u.Email = request.Email
	u.Birthday = request.Birthday
	u.Gender = request.Gender
}

func (u *User) ToRegisterResponse() web.UserRegisterResponse {
	response := web.UserRegisterResponse{}
	response.Name = u.Name
	response.Email = u.Email
	response.Birthday = u.Birthday
	response.Role = u.Role
	response.Status.Name = u.Status.Name
	response.Status.CreatedAt = u.Status.CreatedAt
	response.Status.UpdatedAt = u.Status.UpdatedAt
	return response
}
