package models

import (
	"database/sql"
)

type Mark struct {
	ID        int64         `gorm:"primary_key"`
	CommentId sql.NullInt64 `gorm:"type:bigint REFERENCES comments(id)"`
	Kind      string
	Value     string
	Approval  sql.NullBool
}
