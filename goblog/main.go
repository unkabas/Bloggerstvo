package main

import (
	"github.com/unkabas/Bloggerstvo/internal/mdparser"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type Blog struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Directory   string `json:"directory"`
	Content     string `json:"content,omitempty"`
}

var allBlogs []Blog
var contentDir = "./content"

func main() {
	allBlogs = loadMarkdownPosts(contentDir)

	r := gin.Default()

	// Роуты
	r.GET("/blogs", getAllBlogs)
	r.GET("/blogs/:title", getBlogById)

	r.Run(":8080")
}

func getAllBlogs(c *gin.Context) {
	var blogs []Blog
	for _, b := range allBlogs {
		blogs = append(blogs, Blog{
			ID:          b.ID,
			Title:       b.Title,
			Description: b.Description,
			Directory:   b.Directory,
		})
	}
	c.JSON(http.StatusOK, blogs)
}

func getBlogById(c *gin.Context) {
	title := c.Param("title")
	for _, blog := range allBlogs {
		if strings.EqualFold(blog.Title, title) {
			c.JSON(http.StatusOK, blog)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
}

func loadMarkdownPosts(dir string) []Blog {
	var blogs []Blog
	idCounter := 1

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.EqualFold(filepath.Ext(path), ".md") {
			return nil
		}

		mdContent, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		// Генерация заголовка
		title := generateTitleFromFilename(info.Name())

		blog := Blog{
			ID:          idCounter,
			Title:       title,
			Description: extractDescription(mdContent),
			Directory:   filepath.Dir(path),
			Content:     string(mdparser.MdToHTML(mdContent)), // Используем ваш парсер
		}

		blogs = append(blogs, blog)
		idCounter++
		return nil
	})

	return blogs
}

// Вспомогательные функции
func generateTitleFromFilename(filename string) string {
	name := strings.TrimSuffix(filename, filepath.Ext(filename))
	return strings.Title(strings.ReplaceAll(name, "-", " "))
}

func extractDescription(md []byte) string {
	// Берем первую строку как описание
	firstLine := strings.Split(string(md), "\n")[0]
	return strings.TrimPrefix(firstLine, "# ")
}
