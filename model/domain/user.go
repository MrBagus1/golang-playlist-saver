package domain

import (
	"time"
)

type User struct {
	Id        int       `json:"id" gorm:"primary_key,autoIncrement,not null"`
	Name      string    `json:"name" gorm:"not null"`
	Password  string    `json:"password" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique"`
	Birthday  time.Time `json:"birthday"`
	Gender    string    `json:"gender"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    Status    `json:"status"`
}
