package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dmawardi/go-server/pkg/config"
	"github.com/dmawardi/go-server/pkg/handlers"
	"github.com/dmawardi/go-server/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

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

	// Route handler
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	// The second argument is a handler, but not required due to above
	// returns an error

	fmt.Printf("Starting application on port: %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
