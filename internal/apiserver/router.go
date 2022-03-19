package apiserver

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/usmhk/goserver/pkg/version"
	"net/http"

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
	g.GET("/healthz", func(c *gin.Context) {
		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})
	// version handler
	g.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, version.Get().String())
	})
	pprof.Register(g)

	// swagger api docs
	//g.GET("/swagger/*any", ginswagger.WrapHandler(swaggerFiles.Handler))
}
