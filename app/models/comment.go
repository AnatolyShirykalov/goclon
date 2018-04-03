package models

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID              int64  `gorm:"primary_key" json:"id"`
	ClassicOnlineId string `gorm:"unique_index"`
	Date            *time.Time
	UserID          sql.NullInt64 `gorm:type:bigint REFERENCES users(id)"`
	User            User
	PerformID       sql.NullInt64 `gorm:type:bigint REFERENCES users(id)"`
	Perform         Perform
	Likes           int64
	Text            string `json:"text"`
	Approval        sql.NullBool
}
