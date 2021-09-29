package record

type YoutubeData struct {
	Id          int `gorm:"primary_key,not null"`
	YoutubeLink string
	Title       string
	ChannelId   string
	PublishedAt string
	Description string
}

//PlaylistDetail PlaylistDetail `gorm:"foreignKey: YoutubeId"`
