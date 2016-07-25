package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response for response struct
type Response struct {
	Code    Code        `json:"code,string"`
	Action  interface{} `json:"action,string"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// JSON write data via gin.Context
func JSON(c *gin.Context, response Response) {
	response.Action, _ = c.Get("action")
	c.JSON(http.StatusOK, response)
}

// OK for successful request
func OK(c *gin.Context, data interface{}) {
	JSON(c, Response{
		Code:    StatusOK,
		Message: "ok",
		Data:    data,
	})
}

// Failed for failing request
func Failed(c *gin.Context, code Code, message string) {
	JSON(c, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// ServerError for internal procedure panic
func ServerError(c *gin.Context, message string) {
	JSON(c, Response{
		Code:    StatusInternalServerError,
		Message: message,
		Data:    nil,
	})
}
