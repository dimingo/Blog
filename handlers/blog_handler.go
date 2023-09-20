package handlers

import (
	"example.com/myBlogWorkspace/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"

	"html/template"
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

func GetBlogPostHandler(c *gin.Context) {
	postName := c.Param("postName")

	fmt.Println(postName)

	mdfile, err := os.ReadFile("./markdown/" + postName)

	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusNotFound, "error.tmpl.html", nil)
		c.Abort()
		return
	}
	postHTML := template.HTML(blackfriday.Run([]byte(mdfile)))

	post := models.Post{Title: postName, Content: postHTML}

	c.HTML(http.StatusOK, "post.tmpl.html", gin.H{
		"Title":   post.Title,
		"Content": post.Content,
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
