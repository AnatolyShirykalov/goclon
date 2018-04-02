package models

type User struct {
	ID              int64  `gorm:"primary_key"`
	ClassicOnlineId string `gorm:"unique_index"`
	Name            string
	Password        string
	Comments        []*Comment
}
