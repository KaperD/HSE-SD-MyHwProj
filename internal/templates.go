package myhwproj

import (
	"html/template"
	"net/url"
	"path/filepath"
)

// NewTemplateCache creates cache of html templates
func NewTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.gohtml"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts := template.New(name)

		ts.Funcs(template.FuncMap{
			"isLink": func(s string) bool {
				_, err := url.ParseRequestURI(s)
				return err == nil
			},
		})

		ts, err := ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.gohtml"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
