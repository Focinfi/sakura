package models

import (
	"github.com/Focinfi/sakura/app/i18n"
	"github.com/gin-gonic/gin"
)

const (
	// EmailRegistration for email registration
	EmailRegistration int = 1
	// UserNameRegistration for user_name registration
	UserNameRegistration = 2
	// PhoneRegistration for phone registration
	PhoneRegistration = 3
)

// Context contains data for every request
type Context struct {
	*gin.Context
	Params *RequestParams
}

// RequestParams for http JOSN request body
type RequestParams struct {
	AccessToken string      `json:"access_token"`
	Action      string      `json:"action"`
	ActionToken string      `json:"action_token"`
	UserID      string      `json:"user_id"`
	LoginToken  string      `json:"login_token"`
	Locale      i18n.Locale `json:"locale,string"`

	User             *User  `json:"user,omitempty"`
	RegistrationType int    `json:"registration_type,omitempty"`
	Phone            string `json:"phone,omitempty"`
	Email            string `json:"email,omitempty"`
	UserPassword     string `json:"password"`
	VerificationCode string `json:"verification_code,omitempty"`
}
