package apiserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	ginswagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/usmhk/goserver/pkg/core"
	"github.com/usmhk/goserver/pkg/errno"
)

func loadRouter(g *gin.Engine) {
	installController(g)
}

func installController(g *gin.Engine) {
	// 404 handler
	g.NoRoute(func(c *gin.Context) { core.WriteResponse(c, errno.ErrPageNotFound, nil) })

	// healthz handler.
	g.GET("/healthz",
		func(c *gin.Context) {
			fmt.Println("---------------------------")
			core.WriteResponse(c, nil, map[string]string{"status": "ok"})
		})

	// swagger api docs
	g.GET("/swagger/*any", ginswagger.WrapHandler(swaggerFiles.Handler))
}
