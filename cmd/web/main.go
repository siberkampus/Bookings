package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"udemy/pkg/config"
	"udemy/pkg/handlers"
	"udemy/pkg/renders"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	templateCache, err := renders.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template")
	}
	app.TemplateCache = templateCache
	app.UseCache = false
	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)
	renders.NewTemplates(&app)

	fmt.Printf("Starting application on %s port\n", portNumber)

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	log.Fatal(serve.ListenAndServe())
}
