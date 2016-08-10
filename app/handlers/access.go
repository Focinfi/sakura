package handlers

import (
	"github.com/Focinfi/sakura/app"
	"github.com/Focinfi/sakura/app/errors"
	"github.com/Focinfi/sakura/app/response"
	"github.com/Focinfi/sakura/libs/token"
	"github.com/gin-gonic/gin"
)

// AccessAuth for auth
func AccessAuth(c *gin.Context) {
	requestParams, ok := paramsFromContext(c)
	if !ok {
		response.ServerError(c, "failed to get params from Context")
		c.Abort()
		return
	}

	vals := map[string]string{
		app.NameStr:      app.LoginStr,
		app.UserIDString: requestParams.UserID,
	}

	if requestParams.UserID != "" && token.CheckWithVals(requestParams.LoginToken, vals) {
		response.Failed(c, errors.LoginTokenIsWorng, "login_token is wrong")
		c.Abort()
		return
	}
}
