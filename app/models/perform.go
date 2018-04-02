package models

import (
	"database/sql"
	"time"
)

type Perform struct {
	ID              int64  `gorm:"primary_key"`
	ClassicOnlineId string `gorm:"unique_index"`
	Date            *time.Time
	UserID          sql.NullInt64 `gorm:type:bigint REFERENCES users(id)"`
	User            User
	Likes           int64
	PieceID         sql.NullInt64 `gorm:type:bigint REFERENCES pieces(id)"`
	Piece           Piece
	GroupID         sql.NullInt64 `gorm:type:bigint REFERENCES groups(id)"`
	Group           Group
	Comments        []*Comment
}
