package main

import (
	"fmt"
	"github.com/lazazael/GoWeb5HelloServer/pkg/config"
	"github.com/lazazael/GoWeb5HelloServer/pkg/handlers"
	"github.com/lazazael/GoWeb5HelloServer/pkg/render"
	"log"
	"net/http"
)

const portNumber string = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf(fmt.Sprintf("Starting app on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
