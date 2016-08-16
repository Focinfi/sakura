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

	// common handler
	commonHandler := handlers.NewBase()
	commonHandler.AddHandlerFunc("send_phone_verification_code", handlers.SendPhoneVerificationCode)
	commonHandler.AddHandlerFunc("verify_phone_code", handlers.VerifyPhoneCode)
	api.POST("common", commonHandler.Handle)

	// user handler
	userHandler := handlers.NewBase()
	userHandler.AddHandlerFunc("create", handlers.CreateUser)
	userHandler.AddHandlerFunc("login", handlers.Login)
	api.POST("user", userHandler.Handle)

	public := server.Group("public")
	public.GET("verifyEmail/:token", handlers.VerifyEmail)
}
