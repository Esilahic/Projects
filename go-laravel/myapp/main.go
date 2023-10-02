package main

import celeritas "github.com/Esilahic/Projects/go-laravel"

type application struct {
	App *celeritas.Celeritas
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
