package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cpratap994/bookings-learngo/internal/config"
	"github.com/cpratap994/bookings-learngo/internal/models"
	"github.com/cpratap994/bookings-learngo/internal/render"
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
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {

	tmpmap := make(map[string]string)
	tmpmap["test"] = "Hello test"

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{StringMap: tmpmap})
}

func (repo *Repository) LuxuryRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "luxuryroom.page.tmpl", &models.TemplateData{})
}

func (repo *Repository) PresidentSuite(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "presidentroom.page.tmpl", &models.TemplateData{})
}

func (repo *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

func (repo *Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {
	startDate := r.Form.Get("start")
	endDate := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Dates are %s %s", startDate, endDate)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (repo *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {

	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (repo *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (repo *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}
