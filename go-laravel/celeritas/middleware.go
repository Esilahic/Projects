package celeritas

import "net/http"

func (c *Celeritas) SessionLoad(next http.Handler) http.Handler {
	return c.Sessions.LoadAndSave(next)
}
