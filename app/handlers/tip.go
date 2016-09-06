package handlers

import (
	"github.com/Focinfi/sakura/app/models"
	"github.com/Focinfi/sakura/app/response"
)

// GenerateTip for new action
func GenerateTip(c *models.Context) {
	response.OK(c.Context, "Hello Honey")
}
