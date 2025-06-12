package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Blog struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Directory   string `json:"directory"`
}

var allBlogs []Blog

func getAllBlogs(c *gin.Context) {
	c.JSON(http.StatusOK, allBlogs)
}

func getBlogById(c *gin.Context) {
	title := c.Param("title")
	for _, blog := range allBlogs {
		if blog.Title == title {
			c.JSON(http.StatusOK, blog)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
}

func main() {
	r := gin.Default()

	r.GET("/blogs", getAllBlogs)
	r.GET("/blogs/:title", getBlogById)

	r.Run(":8080")
}
