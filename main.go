package main

import (
	"fmt"
	"net/http"

	"github.com/focinfi/sakura/config"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	mux := http.NewServeMux()
	adminRouter := server.Group("/admin")
	adminRouter.Any("/*w", gin.WrapH(mux))
	config.Routes(server)
	server.Static("/assets", "./public")

	fmt.Printf("Listening on: %v\n", config.Config.Port)
	if err := server.Run(fmt.Sprintf(":%d", config.Config.Port)); err != nil {
		panic(err)
	}
}
