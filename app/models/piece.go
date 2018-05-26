package models

import "database/sql"

type Piece struct {
	ID              int64  `gorm:"primaty_key"`
	ClassicOnlineId string `gorm:"unique_index"`
	Name            string
	ComposerID      sql.NullInt64 `gorm:"type:bigint REFERENCES composers(id)"`
	Composer        Composer
	Performs        []*Perform
}
