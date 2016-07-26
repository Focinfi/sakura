package admin

import (
	"github.com/Focinfi/sakura/db"
	"github.com/qor/admin"
	"github.com/qor/qor"
)

// Admin for data administration
var Admin *admin.Admin

func init() {
	Admin = admin.New(&qor.Config{DB: db.DB})
	Admin.SetSiteName("sakura")
	Admin.SetAuth(AdminUser)

	// Add Dashboard
	Admin.AddMenu(&admin.Menu{Name: "Dashboard", Link: "/admin"})
}
