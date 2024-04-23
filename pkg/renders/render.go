package renders

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"udemy/pkg/config"
	"udemy/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(dt *models.TemplateData)*models.TemplateData{
	return dt
}
func RenderTemplate(w http.ResponseWriter, file string, data *models.TemplateData) {
	var templateCache map[string]*template.Template
	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	t, ok := templateCache[file]
	if !ok {
		log.Fatal("cache not found")
	}
	data=AddDefaultData(data)
	t.ExecuteTemplate(w, "base", data)
}
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./template/*.html")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		tmpl, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		mathes, err := filepath.Glob("./template/*.layout.html")
		if err != nil {
			return myCache, err
		}
		if len(mathes) > 0 {
			tmpl, err = tmpl.ParseGlob("./template/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = tmpl
	}
	return myCache, nil
}
