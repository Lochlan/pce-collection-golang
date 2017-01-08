package main

import (
    "github.com/google/jsonapi"
    "github.com/lochlan/pce-collection-golang/games"
    "net/http"
    "os"
)

var pwd, _ = os.Getwd()
var db = games.InitDB(pwd + "/test.db")

func handler(w http.ResponseWriter, r *http.Request) {
    slug := r.URL.Path[len("/games/"):]
    requested_game := games.ReadGameBySlug(db, slug)
    if requested_game == nil {
        http.Error(w, "404 page not found", 404)
        return
    }

    w.WriteHeader(200)
    w.Header().Set("Content-Type", "application/vnd.api+json")
    if err := jsonapi.MarshalOnePayload(w, requested_game); err != nil {
        http.Error(w, err.Error(), 500)
    }
}

func main() {
    defer db.Close()
    games.CreateTable(db)

    http.HandleFunc("/games/", handler)
    http.ListenAndServe(":8080", nil)
}
