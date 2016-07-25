package models

import (
	"github.com/gin-gonic/gin"
)

// Context contains data for every request
type Context struct {
	*gin.Context
	Params *RequestParams
}

// RequestParams for http JOSN request body
type RequestParams struct {
	AccessToken string `json:"access_token"`
	Action      string `json:"action"`
	ActionToken string `json:"action_token"`
	UserID      string `json:"user_id"`

	User *User `json:"user"`
}
