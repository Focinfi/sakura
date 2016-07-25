package handlers

import (
	"github.com/Focinfi/sakura/app/response"
	"github.com/gin-gonic/gin"
)

// AccessAuth for auth
func AccessAuth(c *gin.Context) {
	requestParams, ok := paramsFromContext(c)
	if !ok {
		response.ServerError(c, "failed to get params from Context")
	}

	if requestParams.AccessToken != "" {
		response.Failed(c, response.AccessTokenIsWrong, "access_token is wrong")
		c.Abort()
		return
	}

	// if requestParams.UserID != ""  && token.CheckLoginToken(token, requestParams.UserID) {
	//
	// }
}
