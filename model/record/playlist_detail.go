package record

type PlaylistDetail struct {
	Id            int `gorm:"not null, autoIncrement"`
	PlaylistId    int
	YoutubeDataId string
	YoutubeData   YoutubeData `gorm:"foreignKey: YoutubeDataId"`
}
