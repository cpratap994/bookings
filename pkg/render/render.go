package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/cpratap994/bookings-learngo/pkg/config"
	"github.com/cpratap994/bookings-learngo/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func Newtemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

//RenderTemplate
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc, _ = CreateTemplateCache()
	} else {
		tc = app.TemplateCache
	}

	parsedTemplate, ok := tc[tmpl]
	if !ok {
		log.Fatal("Failed app")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	_ = parsedTemplate.Execute(buf, td)
	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {

		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}

	return myCache, nil
}
