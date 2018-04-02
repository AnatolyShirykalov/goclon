package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/qor/media"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var DB *gorm.DB
var Config map[string]string

func init() {
	dat, err := ioutil.ReadFile("./config/database.yml")
	if err != nil {
		panic(err)
	}

	Config = make(map[string]string, 0)

	err = yaml.Unmarshal([]byte(dat), &Config)
	if err != nil {
		panic(err)
	}
	log.Printf("connecting to %v, conf: %v", Config["adapter"], Config["settings"])
	DB, err = gorm.Open(Config["adapter"], Config["settings"])
	DB.LogMode(false)
	if err != nil {
		panic(err)
	}
	media.RegisterCallbacks(DB)
	DB.AutoMigrate(&User{}, &Composer{}, &Comment{}, &Group{},
		&GroupPerformer{}, &Perform{}, &Performer{}, &Piece{})
	//&ParserJobArgument{})
	DB.Model(&Piece{}).AddForeignKey("composer_id", "composers(id)", "RESTRICT", "RESTRICT")
	DB.Model(&GroupPerformer{}).AddForeignKey("performer_id", "performers(id)", "RESTRICT", "RESTRICT")
	DB.Model(&GroupPerformer{}).AddForeignKey("group_id", "groups(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Perform{}).AddForeignKey("piece_id", "pieces(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Perform{}).AddForeignKey("group_id", "groups(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Perform{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Comment{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Comment{}).AddForeignKey("perform_id", "performs(id)", "RESTRICT", "RESTRICT")
}
