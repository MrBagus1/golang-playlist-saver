package record

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id        int    `gorm:"primary_key,autoIncrement,not null"`
	Name      string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"unique"`
	Birthday  time.Time
	Gender    string
	Role      string `gorm:"type:enum('ADMIN','USER');default:'USER'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    Status `gorm:"foreignKey:UserId"`
}
