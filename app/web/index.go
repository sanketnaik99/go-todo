package web

import (
	"io/fs"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func WebGroup(router *gin.Engine) * gin.RouterGroup {
	webRouter := router.Group("/web")
	{

		staticFiles, _ := fs.Sub(Files, "assets")
		webRouter.StaticFS("/assets", http.FS(staticFiles))

		webRouter.GET("/", func(c *gin.Context) {
			templ.Handler(HelloForm()).ServeHTTP(c.Writer, c.Request)
		})
	}

	return webRouter
}

