package main

import (
	"fmt"
	"net/http"

	"github.com/Focinfi/sakura/config"
	"github.com/Focinfi/sakura/config/admin"
	"github.com/Focinfi/sakura/config/routes"
	_ "github.com/Focinfi/sakura/db/migrations"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	mux := http.NewServeMux()
	admin.Admin.MountTo("/admin", mux)
	adminRouter := server.Group("/admin")
	adminRouter.Any("/*w", gin.WrapH(mux))
	routes.Routes(server)
	server.Static("/assets", "./public")

	fmt.Printf("Listening on: %v\n", config.Config.Port)
	if err := server.Run(fmt.Sprintf(":%d", config.Config.Port)); err != nil {
		panic(err)
	}
}
