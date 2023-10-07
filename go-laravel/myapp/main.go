package main

import (
	"myapp/handlers"

	celeritas "github.com/Esilahic/Projects/go-laravel"
)

type application struct {
	App      *celeritas.Celeritas
	Handlers *handlers.Handlers
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
