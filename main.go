package main

import (
	"example.com/myBlogWorkspace/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Delims("{{", "}}")
	r.LoadHTMLGlob("./templates/*.tmpl.html")
	// Serve static assets from the "static" directory
	r.Static("/static", "./static")
	routes.SetUpRoutes(r)

	// Start the server
	r.Run(":8080")
}
