package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/media"
)

var DB *gorm.DB
var Config map[string]string

func Init() {
	DB.LogMode(true)
	media.RegisterCallbacks(DB)
	DB.AutoMigrate(&User{}, &Composer{}, &Group{}, &Performer{},
		&GroupPerformer{}, &Piece{}, &Perform{}, &Comment{}, &Mark{}, &SiteUser{})
}
