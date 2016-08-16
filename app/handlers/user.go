package handlers

import (
	"github.com/Focinfi/sakura/app/i18n"
	"github.com/Focinfi/sakura/app/log"
	"github.com/Focinfi/sakura/app/models"
	"github.com/Focinfi/sakura/app/response"
	"github.com/Focinfi/sakura/app/services"
	"github.com/Focinfi/sakura/db"
	"github.com/Focinfi/sakura/libs/token"
	"github.com/gin-gonic/gin"
)

// NewUser for new action
func NewUser(c *models.Context) {

}

// CreateUser handles create action
func CreateUser(c *models.Context) {
	if err := services.CreateUser(c.Params); err != nil {
		response.Failed(c.Context, err.Code(), err.Message(c.Params.Locale))
		return
	}

	response.OK(c.Context, nil)
}

// VerifyEmail handlers verify_email action
func VerifyEmail(c *gin.Context) {
	// TODO: replace string with HTML
	tkn := c.Query("token")

	if !token.CheckSimple(tkn, "email_verification") {
		c.String(200, i18n.T(i18n.Locale(c.Param("locale")), "failed_to_verify_email"))
		return
	}

	email := token.GetParam(tkn, "email")
	// find the user
	user := &models.User{Email: email}
	query := db.DB.Where(user).First(user)
	if query.Error != nil {
		log.DBError(query.Value, query.Error, "failed to get user")
		return
	}

	if user.ID == 0 {
		c.String(200, i18n.T(i18n.Locale(c.Param("locale")), "email_does_not_exist"))
		return
	}

	if user.Verified {
		c.String(200, i18n.T(i18n.Locale(c.Param("locale")), "email_verified"))
		return
	}

	update := db.DB.Model(user).UpdateColumn("verified", true)
	if update.Error != nil {
		log.DBError(update.Value, update.Error, "failed_to_verify_user")
		return
	}

	c.String(200, i18n.T(i18n.Locale(c.Param("locale")), "email_verified"))
}
