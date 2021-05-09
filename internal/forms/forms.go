package forms

import (
	"net/http"
	"net/url"
	"strings"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {

	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {

	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This filed should not be empty")
		}
	}
}

//Has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {

	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "this field can not be blank")
		return false
	}

	return true
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
