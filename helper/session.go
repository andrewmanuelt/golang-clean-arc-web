package helper

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

type SessionInterface interface {
	SetSession(w http.ResponseWriter, r *http.Request, key string, value string)
	GetSession(w http.ResponseWriter, r *http.Request, key string) string
	DeleteAllSession(w http.ResponseWriter, r *http.Request)
}

var (
	store = sessions.NewCookieStore([]byte("SESSION_KEY"))
)

type sessionInterfaceImpl struct{}

func (*sessionInterfaceImpl) DeleteAllSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Options.MaxAge = -1

	session.Save(r, w)
}

func (*sessionInterfaceImpl) GetSession(w http.ResponseWriter, r *http.Request, key string) string {
	session, _ := store.Get(r, "session")

	if len(session.Values) == 0 {
		return ""
	}

	x := session.Values[key]

	return fmt.Sprintf("%s", x)
}

func (*sessionInterfaceImpl) SetSession(w http.ResponseWriter, r *http.Request, key string, value string) {
	session, _ := store.Get(r, "session")

	session.Values[key] = value

	err := session.Save(r, w)

	ShowError(err)
}

func NewSession() SessionInterface {
	return &sessionInterfaceImpl{}
}
