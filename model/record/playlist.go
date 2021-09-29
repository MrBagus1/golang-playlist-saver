package record

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

type Playlist struct {
	Id             int `gorm:"not null, autoIncrement"`
	Name           string
	UserId         int `gorm:"foreignKey: Id"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      null.Time
	PlaylistDetail []PlaylistDetail `gorm:"foreignKey: PlaylistId"`
}
