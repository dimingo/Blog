package routes

import (
	"example.com/myBlogWorkspace/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {

	router.GET("/", handlers.Welcome)
	router.POST("/posts", handlers.CreateBlogPostHandler)
	router.GET("/:postName", handlers.GetBlogPostHandler)
	/*
		router.GET("/posts", handlers.ListBlogPostsHandler)
		router.GET("/posts/:id", handlers.GetBlogPostHandler)
		router.PUT("/posts/:id", handlers.UpdateBlogPostHandler)
		router.DELETE("/posts/:id", handlers.DeleteBlogPostHandler)
	*/

}
