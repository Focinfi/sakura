package handlers

import (
	"github.com/Focinfi/sakura/app/response"
	"github.com/focinfi/sakura/app/models"
	"github.com/gin-gonic/gin"
)

// ParseParams parses params
func ParseParams(c *gin.Context) {
	// Parsing Params
	params := &models.RequestParams{}
	if err := c.Bind(params); err != nil {
		response.Failed(c, response.JSONBodyParsingError, "failed to parsing JOSN boday")
		c.Abort()
		return
	}

	c.Set("params", params)
}

func paramsFromContext(c *gin.Context) (*models.RequestParams, bool) {
	params, has := c.Get("params")
	requestParams, ok := params.(*models.RequestParams)
	return requestParams, !has || !ok
}