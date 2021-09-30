package record

type PlaylistDetail struct {
	Id            int `gorm:"not null, autoIncrement"`
	PlaylistId    int
	YoutubeDataId string      `gorm:"constraint:OnDelete:CASCADE"`
	YoutubeData   YoutubeData `gorm:"foreignKey: YoutubeDataId;constraint:OnDelete:CASCADE"`
}
//