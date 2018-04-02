package admin

import (
	"../../app/models"
	"github.com/qor/admin"
	"github.com/qor/qor"
)

var Admin *admin.Admin

func init() {
	Admin = admin.New(&qor.Config{DB: models.DB})
	Admin.SetSiteName("Classic-online Parser")

	Admin.AddResource(&models.User{})
	Admin.AddResource(&models.Comment{})
	Admin.AddResource(&models.Composer{})
	Admin.AddResource(&models.Group{})
	Admin.AddResource(&models.Perform{})
	Admin.AddResource(&models.Performer{})
	Admin.AddResource(&models.Piece{})
}
