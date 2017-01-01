package main

import (
    "fmt"
    "github.com/lochlan/pce-collection-golang/games"
    "net/http"
    "os"
)

var pwd, _ = os.Getwd()
var db = games.InitDB(pwd + "/test.db")

func handler(w http.ResponseWriter, r *http.Request) {
    slug := r.URL.Path[1:]
    requested_game := games.ReadGameBySlug(db, slug)
    if requested_game == nil {
        return
    }
    fmt.Fprintf(w, "<p>" + requested_game.ToString() + "</p>")
    fmt.Fprintf(w, "<p>" + requested_game.Developer + "</p>")
}

func main() {
    defer db.Close()
    games.CreateTable(db)

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
