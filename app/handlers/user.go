package handlers

import (
	"github.com/Focinfi/sakura/app/models"
	"github.com/Focinfi/sakura/app/response"
	"github.com/Focinfi/sakura/app/services"
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
