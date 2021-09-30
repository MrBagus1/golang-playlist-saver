package record

import (
	"gorm.io/gorm"
	"time"
)

type Playlist struct {
	Id             int `gorm:"not null, autoIncrement"`
	Name           string
	UserId         int `gorm:"foreignKey: Id"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	PlaylistDetail []PlaylistDetail `gorm:"foreignKey: PlaylistId;constraint:OnDelete:CASCADE"`
}
