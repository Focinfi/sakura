package handlers

import (
	"github.com/Focinfi/sakura/app/models"
	"github.com/Focinfi/sakura/app/response"
	"github.com/Focinfi/sakura/app/services"
)

// NewUser for new action
func NewUser(c *models.Context) {

}

// CreateUser for create action
func CreateUser(c *models.Context) {
	rType := c.Params.RegistrationType
	user := c.Params.User

	if err := services.CreateUser(rType, user); err != nil {
		response.Failed(c.Context, err.Code(), err.Message(c.Params.Locale))
		return
	}

	response.OK(c.Context, nil)
}
