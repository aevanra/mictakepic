package session

import (
    "os"
    "net/http"
    "github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

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