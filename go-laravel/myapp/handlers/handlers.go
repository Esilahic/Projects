package handlers

import (
	"net/http"

	"github.com/CloudyKit/jet/v6"
	celeritas "github.com/Esilahic/Projects/go-laravel"
)

type Handlers struct {
	App *celeritas.Celeritas
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering", err)
	}
}

func (h *Handlers) GoPage(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.GoPage(w, r, "home", nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering", err)
	}
}

func (h *Handlers) JetPage(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.JetPage(w, r, "jet-template", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering", err)
	}
}

func (h *Handlers) SessionTest(w http.ResponseWriter, r *http.Request) {
	myData := "bar"

	h.App.Sessions.Put(r.Context(), "foo", myData)

	myValue := h.App.Sessions.GetString(r.Context(), "foo")

	vars := make(jet.VarMap)
	vars.Set("foo", myValue)

	err := h.App.Render.JetPage(w, r, "sessions", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering", err)
	}
}
