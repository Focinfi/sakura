package handlers

import (
	"github.com/Focinfi/sakura/app/models"
	"github.com/Focinfi/sakura/app/response"
	"github.com/gin-gonic/gin"
)

// Handler handles params and writes response into c.
type Handler interface {
	Handle(c *models.Context)
}

// HandlerFunc defines Handle
type HandlerFunc func(c *models.Context)

// Handle impelements Handler
func (f HandlerFunc) Handle(c *models.Context) {
	f.Handle(c)
}

// Base for basic handling
type Base struct {
	Handlers map[string]Handler
}

// NewBase allocates and returns a new Base
func NewBase() *Base {
	return &Base{Handlers: map[string]Handler{}}
}

// AddHandler add routes for base.routes
func (base *Base) AddHandler(name string, handler Handler) {
	base.Handlers[name] = handler
}

// AddHandlerFunc add routes for base.routes
func (base *Base) AddHandlerFunc(name string, f func(c *models.Context)) {
	base.Handlers[name] = HandlerFunc(f)
}

// Handle handles request
func (base *Base) Handle(c *gin.Context) {
	requestParams, ok := paramsFromContext(c)
	if !ok {
		response.ServerError(c, "failed to get params from Context")
	}

	// Dispatch task
	handler, ok := base.Handlers[requestParams.Action]
	if !ok {
		response.Failed(c, response.ActionIsNotAllowed, "")
		c.Abort()
		return
	}

	handler.Handle(&models.Context{Context: c, Params: requestParams})
}
