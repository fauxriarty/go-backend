package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fauxriarty/go-backend/pkg/config"
	"github.com/fauxriarty/go-backend/pkg/handlers"
	"github.com/fauxriarty/go-backend/pkg/render"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	// its false bec we're in development mode since we want to see realtime changes not the cached version
	app.UseCache = false // set to true in production for efficiency

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s ", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	_ = srv.ListenAndServe()

}
