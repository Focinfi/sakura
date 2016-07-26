package routes

import (
	"github.com/Focinfi/sakura/app/handlers"
	"github.com/gin-gonic/gin"
)

// Routes configure routes
func Routes(server *gin.Engine) {
	// middlewares
	api := server.Group("api",
		handlers.ParseParams,
		handlers.AccessAuth,
	)

	// user handler
	userHandler := handlers.NewBase()
	userHandler.AddHandlerFunc("create", handlers.CreateUser)

	api.POST("user", userHandler.Handle)
}
