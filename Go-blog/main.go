package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/unkabas/Bloggerstvo/mdParser" // замените на ваш реальный модуль
)

const (
	contentDir = "./content" // Директория с Markdown файлами
	baseURL    = "/blogs"    // Базовый URL путь
)

type BlogPost struct {
	FilePath string
	URLPath  string
	HTML     []byte
}

func processMarkdownFiles() ([]BlogPost, error) {
	var posts []BlogPost

	err := filepath.Walk(contentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Пропускаем директории и не-Markdown файлы
		if info.IsDir() || strings.ToLower(filepath.Ext(path)) != ".md" {
			return nil
		}

		// Читаем файл
		file, err := os.Open(path)
		if err != nil {
			log.Printf("Ошибка открытия файла %s: %v", path, err)
			return nil
		}
		defer file.Close()

		mdContent, err := io.ReadAll(file)
		if err != nil {
			log.Printf("Ошибка чтения файла %s: %v", path, err)
			return nil
		}

		// Конвертируем в HTML
		htmlContent := mdParser.MdToHTML(mdContent)

		// Формируем URL путь
		relPath, _ := filepath.Rel(contentDir, path)
		urlPath := filepath.Join(baseURL, strings.TrimSuffix(relPath, filepath.Ext(relPath)))

		posts = append(posts, BlogPost{
			FilePath: path,
			URLPath:  urlPath,
			HTML:     htmlContent,
		})

		return nil
	})

	return posts, err
}

func main() {
	posts, err := processMarkdownFiles()
	if err != nil {
		log.Fatalf("Ошибка обработки файлов: %v", err)
	}

	// Выводим результаты
	for _, post := range posts {
		fmt.Printf("\n=== Файл: %s ===\n", post.FilePath)
		fmt.Printf("URL: %s\n", post.URLPath)
		fmt.Printf("Контент:\n%s\n", string(post.HTML))
	}

	fmt.Printf("\nОбработано %d статей\n", len(posts))
}
