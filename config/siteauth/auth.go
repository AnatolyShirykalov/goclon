package siteauth

import (
	"../../app/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/qor/auth"
	"github.com/qor/auth/auth_identity"
	"github.com/qor/auth/providers/password"
	"github.com/qor/redirect_back"
	"github.com/qor/session/manager"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

var Auth *auth.Auth

func Init() {
	dat, err := ioutil.ReadFile("./config/auth.yml")
	if err != nil {
		panic(err)
	}
	var conf map[string]string
	err = yaml.Unmarshal([]byte(dat), &conf)
	if err != nil {
		panic(err)
	}
	ss, ok := conf["signed_string"]
	if !ok {
		panic("Need signed_string in config/auth.yml")
	}
	Auth = auth.New(&auth.Config{
		DB:        models.DB,
		UserModel: models.SiteUser{},

		SessionStorer: &auth.SessionStorer{
			SessionName:    "_rm_auth",
			SessionManager: manager.SessionManager,
			SigningMethod:  jwt.SigningMethodHS256,
			SignedString:   ss,
		},
		Redirector: auth.Redirector{redirect_back.New(&redirect_back.Config{
			SessionManager:  manager.SessionManager,
			IgnoredPrefixes: []string{"/auth"},
		})}})

	models.DB.AutoMigrate(&auth_identity.AuthIdentity{})

	Auth.RegisterProvider(password.New(&password.Config{}))
}
