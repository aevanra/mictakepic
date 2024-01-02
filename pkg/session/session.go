package session

import (
    "encoding/gob"
    "net/http"
    "github.com/gorilla/sessions"
    "github.com/gorilla/securecookie"
    "github.com/aevanra/mictakepic/pkg/objects"
)

var store *sessions.CookieStore

func init() {
    gob.Register(&obj.User{})

    sess_key := securecookie.GenerateRandomKey(64)
    hash_key := securecookie.GenerateRandomKey(32)

    store = sessions.NewCookieStore(sess_key, hash_key)
    store.Options = &sessions.Options{
        Path:     "/",
        MaxAge:   300,
        HttpOnly: true,
    }
}

func Get(req *http.Request) (*sessions.Session, error) {
    return store.Get(req, "default-session-name")
}

func GetNamed(req *http.Request, name string) (*sessions.Session, error) {
    return store.Get(req, name)
}

func SetSessionValue(r *http.Request, w http.ResponseWriter, key string, value interface{}) error {
    ses, err := Get(r)
    if err != nil {
        return err
    }

    ses.Values[key] = value
    err = ses.Save(r, w)
    if err != nil {
        return err
    }

    return nil
}

func GetSessionValue(r *http.Request, key string) interface{} {
    ses, _ := Get(r)
    val := ses.Values[key]

    return val
}

func ClearSession(r *http.Request, w http.ResponseWriter) error {
    ses, err := Get(r)
    if err != nil {
        return err
    }

    ses.Options.MaxAge = -1
    err = ses.Save(r, w)
    if err != nil {
        return err
    }

    return nil
}
