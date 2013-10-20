package main

import (
    "encoding/json"
    "net/http"
    "fmt"
    "github.com/yasuoza/msgcookie"
    "io/ioutil"
)

var (
    cookieName string
    secret     string
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie(cookieName)
    if err != nil && err != http.ErrNoCookie {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    v, err := msgcookie.Decode(cookie, secret)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if s, ok := v["session_id"].([]uint8); !ok {
        http.Error(w, "invalid_session_id", http.StatusInternalServerError)
        return
    } else {
        fmt.Fprintf(w, "session id:%q, visits:%d", string(s), v["visits"])
        return
    }
    fmt.Fprint(w, "Bake cookie first!")
}

func main() {
    var config struct {
        CookieName string `json:"cookie_name"`
        Secret     string `json:"cookie_secret"`
    }
    configData, err := ioutil.ReadFile("config.json")
    if err != nil {
        fmt.Println("err", err)
    }
    err = json.Unmarshal(configData, &config)
    if err != nil {
        panic(err)
    }
    cookieName = config.CookieName
    secret = config.Secret
    http.HandleFunc("/", rootHandler)
    http.ListenAndServe(":8080", nil)
}
