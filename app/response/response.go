package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response for response struct
type Response struct {
	Code    int         `json:"code,string"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// JSON write data via gin.Context
func JSON(c *gin.Context, response Response) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, response)
}

// OK for successful request
func OK(c *gin.Context, data interface{}) {
	JSON(c, Response{
		Code:    200,
		Message: "ok",
		Data:    data,
	})
}

// Failed for failing request
func Failed(c *gin.Context, code int, message string) {
	JSON(c, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// ServerError for internal procedure panic
func ServerError(c *gin.Context, message string) {
	JSON(c, Response{
		Code:    500,
		Message: message,
		Data:    nil,
	})
}
