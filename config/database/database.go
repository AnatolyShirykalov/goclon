package database

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"../..//app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	yaml "gopkg.in/yaml.v2"
)

// Database config
type Database struct {
	Adapter  *string
	Host     *string
	Encoding *string
	Port     *int
	Pool     *int
	Username *string
	Password *string
	Db       *string `yaml:"database"`
}

// DB Gorm DB
var DB *gorm.DB

// Config database config data
var Config Database

// Configs possible DB configs
var Configs map[string]Database

func Init() {
	var err error

	dat, err := ioutil.ReadFile("./config/database.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(dat), &Configs)
	if err != nil {
		panic(err)
	}

	env := os.Getenv("RAILS_ENV")
	if env == "" {
		env = "development"
	}
	Config = Configs[env]
	//spew.Dump(Config)

	if *Config.Adapter == "postgresql" {
		connstr := make([]string, 0)
		if Config.Host != nil {
			connstr = append(connstr, "host="+*Config.Host)
		}
		if Config.Port != nil {
			connstr = append(connstr, "port="+strconv.Itoa(*Config.Port))
		}
		if Config.Username != nil {
			connstr = append(connstr, "user="+*Config.Username)
		}
		if Config.Db != nil {
			connstr = append(connstr, "dbname="+*Config.Db)
		}
		if Config.Password != nil {
			connstr = append(connstr, "password="+*Config.Password)
		}
		connstr = append(connstr, "sslmode=disable")
		cs := strings.Join(connstr, " ")
		log.Printf("connecting to pg, conf: %v", cs)
		DB, err = gorm.Open("postgres", cs)
		if err != nil {
			panic(err)
		}
	} else {
		panic("Unknown DB adapter: " + *Config.Adapter)
	}
	//log.Println("enable log mode")
	//DB.LogMode(true)
	DB.DB().SetMaxOpenConns(20)

	models.DB = DB
}
