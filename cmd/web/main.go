package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cpratap994/bookings-learngo/pkg/config"
	"github.com/cpratap994/bookings-learngo/pkg/handlers"
	"github.com/cpratap994/bookings-learngo/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("app failed")
	}

	app.TemplateCache = tc
	app.UseCache = true

	render.Newtemplate(&app)
	// repo settings
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	///http.HandleFunc("/", handlers.Repo.Home)
	///http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting app at port# " + portNumber)
	///http.ListenAndServe(portNumber, nil)
	srv := &http.Server{

		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	log.Fatal(err)
}
