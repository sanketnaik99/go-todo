package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Value string `json:"value"`
}

func ApiGroup(router *gin.Engine) * gin.RouterGroup {
	
	apiRouter := router.Group("/api")
	{

		apiRouter.POST("/todo",func(c *gin.Context) {
			value := c.PostForm("value") 
	
			html := fmt.Sprintf(`
				<li class="flex items-center bg-gray-50 p-2 rounded">
				<input type="checkbox" class="form-checkbox h-5 w-5 text-blue-600 rounded focus:ring-blue-500">
				<span class="ml-4 flex-grow text-gray-500">%s</span> 
				</li>`, value)

			// Set the Content-Type header to "text/html"
			c.Header("Content-Type", "text/html") 

			// Send the HTML string as the response
			c.String(http.StatusOK, html) 
		})
	}

	return apiRouter
}