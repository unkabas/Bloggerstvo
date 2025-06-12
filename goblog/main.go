package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/unkabas/Bloggerstvo/internal/mdparser"
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
	// Проверка существования папки content
	contentPath := "/Users/unkabas/Desktop/goblog/content"
	if _, err := os.Stat(contentPath); os.IsNotExist(err) {
		log.Fatalf("Папка 'content' не найдена по пути: %s", contentPath)
	}

	// Загрузка markdown файлов
	var err error
	allBlogs, err = loadMarkdownPosts(contentPath)
	if err != nil {
		log.Fatalf("Ошибка загрузки markdown файлов: %v", err)
	}
	log.Printf("Загружено %d статей", len(allBlogs))

	r := gin.Default()

	// Роуты
	r.GET("/blogs", getAllBlogs)
	r.GET("/blogs/:title", getBlogByTitle)

	log.Println("Сервер запущен на http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
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

func getBlogByTitle(c *gin.Context) {
	title := c.Param("title")
	for _, blog := range allBlogs {
		if strings.EqualFold(blog.Title, title) {
			c.JSON(http.StatusOK, blog)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
}

func getWorkDir() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Ошибка получения рабочей директории: %v", err)
	}
	return wd
}

func loadMarkdownPosts(dir string) ([]Blog, error) {
	var blogs []Blog
	idCounter := 1

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.EqualFold(filepath.Ext(path), ".md") {
			return nil
		}

		mdContent, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Ошибка чтения файла %s: %v", path, err)
			return nil
		}

		blog := Blog{
			ID:          idCounter,
			Title:       generateTitleFromFilename(info.Name()),
			Description: extractDescription(mdContent),
			Directory:   filepath.Dir(path),
			Content:     string(mdparser.MdToHTML(mdContent)),
		}

		blogs = append(blogs, blog)
		idCounter++
		return nil
	})

	return blogs, err
}

func generateTitleFromFilename(filename string) string {
	name := strings.TrimSuffix(filename, filepath.Ext(filename))
	name = strings.ReplaceAll(name, "-", " ")
	name = strings.ReplaceAll(name, "_", " ")
	return strings.Title(name)
}

func extractDescription(md []byte) string {
	lines := strings.Split(string(md), "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" && !strings.HasPrefix(trimmed, "#") {
			return trimmed
		}
	}
	return "No description available"
}
