package engine

import (
	"embed"
	"net/http"

	"github.com/gofiber/template/html/v2"
)

var (
	//go:embed templates/*
	templatesfs embed.FS
)

func InitViewEngine() *html.Engine {
	engine := html.NewFileSystem(http.FS(templatesfs), ".html")
	return engine
}
