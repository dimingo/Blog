package handlers

import (
	"example.com/myBlogWorkspace/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func Welcome(c *gin.Context) {
	var posts []string

	files, err := os.ReadDir("./markdown")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		posts = append(posts, file.Name())
	}

	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"posts": posts,
	})
}

func CreateBlogPostHandler(c *gin.Context) {
	var newPost models.BlogPost
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the new blog post to your data storage (e.g., a database or slice).
	// You may also want to validate the input data.

	// Return a success response.
	c.JSON(http.StatusCreated, newPost)
}
