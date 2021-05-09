package models

import "github.com/cpratap994/bookings-learngo/internal/forms"

type TemplateData struct {
	StringMap map[string]string
	CSRFToken string
	Data      map[string]interface{}
	Form      *forms.Form
	Flash     string
	Warning   string
	Error     string
}
