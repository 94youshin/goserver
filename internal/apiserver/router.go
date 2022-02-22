package apiserver

import "github.com/gin-gonic/gin"

func loadRouter(g *gin.Engine) {
	installController(g)
}

func installController(g *gin.Engine) {
	g.NoRoute(func(c *gin.Context) { c.String(200, "page not found.") })
}
