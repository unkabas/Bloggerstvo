package mdParser

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"os"
)

// Convert Markdown to HTML with full options
func MdToHTML(md []byte) []byte {
	// Configure parser extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// Configure HTML renderer
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

// Optional: Print AST tree for debugging
func PrintAST(md []byte) {
	p := parser.New()
	doc := p.Parse(md)
	ast.Print(os.Stdout, doc)
}
