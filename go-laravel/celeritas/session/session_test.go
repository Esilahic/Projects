package session

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alexedwards/scs/v2"
)

func TestSession_InitSession(t *testing.T) {
	c := &Session{
		CookieLifetime: "100",
		CookiePersist:  "true",
		CookieName:     "Celeritas",
		CookieDomain:   "localhost",
		SessionType:    "cookie",
	}

	var sm *scs.SessionManager

	ses := c.InitSession()

	var sesKind reflect.Kind
	var sesType reflect.Type

	rv := reflect.ValueOf(ses)

	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		fmt.Println("for loop", rv.Kind(), rv.Type(), rv)
		sesKind = rv.Kind()
		sesType = rv.Type()

		rv = rv.Elem()
	}

	if !rv.IsValid() {
		t.Error("invalid kind or type; kind:", rv.Kind(), "type:", rv.Type())
	}

	if sesKind != reflect.ValueOf(sm).Kind() {
		t.Error("wrong kind returned testing cookie session. Expected:", reflect.ValueOf(sm).Kind(), "got", sesKind)
	}

	if sesType != reflect.ValueOf(sm).Type() {
		t.Error("wrong type returned testing cookie session. Expected:", reflect.ValueOf(sm).Type(), "got", sesType)
	}
}
