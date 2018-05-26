package models

import "database/sql"

type GroupPerformer struct {
	ID          int64         `gorm:"primary_key"`
	PerformerID sql.NullInt64 `gorm:"type:bigint REFERENCES performers(id)"`
	Performer   Performer
	GroupID     sql.NullInt64 `gorm:"type:bigint REFERENCES groups(id)"`
	Group       Group
}
