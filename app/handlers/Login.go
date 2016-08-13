package handlers

import (
	"github.com/Focinfi/sakura/app/errors"
	"github.com/Focinfi/sakura/app/i18n"
	"github.com/Focinfi/sakura/app/models"
	"github.com/Focinfi/sakura/app/response"
	"github.com/Focinfi/sakura/db"
	"github.com/Focinfi/sakura/libs/utils"
)

// Login handles login action
func Login(c *models.Context) {
	var user *models.User
	if c.Params.Phone != "" {
		user = &models.User{Phone: c.Params.Phone}
	} else {
		user = &models.User{Email: c.Params.Email}
	}

	if err := db.DB.First(user).Error; err != nil {
		response.ServerError(c.Context, "failed to query db")
		return
	}

	if user.ID == 0 {
		response.Failed(c.Context,
			errors.EmailIsNonexistent,
			i18n.T(c.Params.Locale, "email_is_nonexistent"))
		return
	}

	if user.Password != utils.Sha256(c.Params.UserPassword+user.Solt) {
		response.Failed(c.Context,
			errors.PasswordIsWrong, i18n.T(c.Params.Locale, "password_is_wrong"))
		return
	}

	response.OK(c.Context, nil)
}
