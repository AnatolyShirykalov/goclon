package models

type Group struct {
	ID              int64 `gorm:"primary_key"`
	GroupPerformers []*GroupPerformer
}
