package models

type Composer struct {
	ID              int64  `gorm:"primaty_key"`
	ClassicOnlineId string `gorm:"unique_index"`
	Name            string
	Pieces          []*Piece
}
