package handlers

import (
	"strconv"

	"github.com/Focinfi/sakura/app/errors"
	"github.com/Focinfi/sakura/app/i18n"
	"github.com/Focinfi/sakura/app/log"
	"github.com/Focinfi/sakura/app/models"
	"github.com/Focinfi/sakura/app/response"
	"github.com/Focinfi/sakura/db"
	"github.com/Focinfi/sakura/libs/token"
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

	query := db.DB.First(user)
	if query.Error != nil {
		log.DBError(query.Value, query.Error, "failed to get user")
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

	tkn, err := token.New("login", 3600*24*30).Set("user_id", strconv.Itoa(int(user.ID))).Sign()
	if err != nil {
		log.LibError("token", "failed to create a login token")
		response.ServerError(c.Context, "failed to make login token")
		return
	}
	response.OK(c.Context, response.Login{Token: tkn})
}
