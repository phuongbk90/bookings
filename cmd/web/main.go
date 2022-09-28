package main

import (
	"fmt"
	"github.com/phuongbk90/bookings/pkg/config"
	"github.com/phuongbk90/bookings/pkg/handlers"
	"github.com/phuongbk90/bookings/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println("Starting application on port: ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
