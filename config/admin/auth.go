package admin

import (
	"github.com/Focinfi/sakura/app/models"
	"github.com/qor/admin"
	"github.com/qor/qor"
)

// Auth for authentification
type Auth struct {
	Name     string `env:"ADMIN_NAME"`
	Password string `env:"ADMIN_PASSWORD"`
	Phone    string `env:"ADMIN_PHONE"`
}

// AdminUser is user with admin Permission
var AdminUser = &Auth{}

// DisplayName implements qor.CurrentUser interface
func (auth *Auth) DisplayName() string {
	return auth.Name
}

// LoginURL returns login url
func (*Auth) LoginURL(c *admin.Context) string {
	return "/admin"
}

// LogoutURL returns logout url,
// right now just for testing
func (*Auth) LogoutURL(c *admin.Context) string {
	return "/admin"
}

// GetCurrentUser returns current user,
// right now just for testing
func (*Auth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	return &models.User{Name: "Admin"}
}
