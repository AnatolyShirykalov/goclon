package admin

import (
	"../../app/models"
	auth "../siteauth"
	"fmt"
	"github.com/qor/admin"
	"github.com/qor/qor"
)

var Admin *admin.Admin

type AdminAuth struct{}

func (AdminAuth) LoginURL(c *admin.Context) string {
	return "/auth/login"
}

func (AdminAuth) LogoutURL(c *admin.Context) string {
	return "/auth/logout"
}

func (AdminAuth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	currentUser := auth.Auth.GetCurrentUser(c.Request)
	if currentUser != nil {
		qorCurrentUser, ok := currentUser.(qor.CurrentUser)
		if !ok {
			fmt.Printf("User %#v haven't implement qor.CurrentUser interface\n", currentUser)
		}
		return qorCurrentUser
	}
	return nil
}
func Init() {
	Admin = admin.New(&admin.AdminConfig{
		DB:   models.DB,
		Auth: &AdminAuth{},
	})
	Admin.SetSiteName("Classic-online Parser")

	Admin.AddResource(&models.User{})
	Admin.AddResource(&models.Comment{})
	Admin.AddResource(&models.Composer{})
	Admin.AddResource(&models.Group{})
	Admin.AddResource(&models.Perform{})
	Admin.AddResource(&models.Performer{})
	Admin.AddResource(&models.Piece{})
}
