package main

import (
	"log"
	"myapp/handlers"
	"os"

	celeritas "github.com/Esilahic/Projects/go-laravel"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//init Celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "myapp"

	MyHandlers := &handlers.Handlers{
		App: cel,
	}

	app := &application{
		App:      cel,
		Handlers: MyHandlers,
	}

	app.App.Routes = app.routes()

	return app
}
