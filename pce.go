package main

import (
    "fmt"
    "net/http"
    "./games"
)

func handler(w http.ResponseWriter, r *http.Request) {
    gradius := &games.PCEngineGame{
        Game: games.Game{
            Name: "Gradius",
        },
        Ean13: "4 988602 585032",
    }
    fmt.Fprintf(w, gradius.ToString())
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
