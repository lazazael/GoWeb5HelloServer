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

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	//http.HandleFunc("/divide", handlers.Divide)

	/*
		http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
			n, err := fmt.Fprintf(w, "Hello,world!")
			fmt.Println(fmt.Sprintf("Number of bytes written: %d",n))
			if err != nil{
				fmt.Println(err)
			}
		})
	*/

	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
