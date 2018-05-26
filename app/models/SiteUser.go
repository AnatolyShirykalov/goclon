package models

import (
	"database/sql"
)

type SiteUser struct {
	ID    int64 `gorm:"primary_key"`
	Name  sql.NullString
	Login sql.NullString `gorm:"unique_index"`
}

func (su *SiteUser) DisplayName() string {
	return su.Name.String
}
