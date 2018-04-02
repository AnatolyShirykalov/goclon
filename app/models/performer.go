package models

type Performer struct {
	ID              int64  `gorm:"primary_key"`
	ClassicOnlineId string `gorm:"unique_index"`
	Instrument      string
	Url             string
	GroupPerformers []*GroupPerformer
	Name            string
}
