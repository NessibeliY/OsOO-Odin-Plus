package templates

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

type Tmpl struct{}

func NewTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	cwd = trimCwd(cwd)

	pages, err := filepath.Glob(cwd + "/ui/html/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

func trimCwd(cwd string) string {
	lastIndex := strings.LastIndex(cwd, "exchange-rates-client")

	return cwd[:lastIndex+len("exchange-rates-client")]
}
