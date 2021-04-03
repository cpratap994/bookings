package handlers

import (
	"net/http"

	"github.com/cpratap994/bookings-learngo/pkg/config"
	"github.com/cpratap994/bookings-learngo/pkg/models"
	"github.com/cpratap994/bookings-learngo/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {

	tmpmap := make(map[string]string)
	tmpmap["test"] = "Hello test"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: tmpmap})
}
