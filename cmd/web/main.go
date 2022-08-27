package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/dmawardi/go-server/pkg/config"
	"github.com/dmawardi/go-server/pkg/handlers"
	"github.com/dmawardi/go-server/pkg/render"
)

// Init state (incl. templates)
var app config.AppConfig

const portNumber = ":8080"

var session *scs.SessionManager

func main() {
	// Change this to true in production
	app.InProduction = false

	session = scs.New()
	// Set session lifetime to 24 hours
	session.Lifetime = 24 * time.Hour
	// Persist when browser restarted
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	// For https set to true
	session.Cookie.Secure = app.InProduction
	// Set session to state variable: Session
	app.Session = session

	// create cache
	createdCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Failed to create template cache")
	}
	// Store created cache in app config
	app.TemplateCache = createdCache
	// Config: Use cache
	// if false, will serve updated file (dev mode)
	app.UseCache = false

	// Create new handler repository
	repo := handlers.NewRepo(&app)
	// Build new handlers
	handlers.UpdateRepositoryHandlers(repo)

	// Sets template cache for render package
	render.SetTemplate(&app)

	fmt.Printf("Starting application on port: %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
