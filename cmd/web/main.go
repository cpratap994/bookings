package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/cpratap994/bookings-learngo/internal/config"
	"github.com/cpratap994/bookings-learngo/internal/handlers"
	"github.com/cpratap994/bookings-learngo/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	//var app config.AppConfig

	// change this to true when in production
	app.InProduction = false

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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
