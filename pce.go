package main

import (
    "fmt"
    "net/http"
)

type Game struct {
    Name string
    Genre string
}

func (g *Game) ToString() string {
    return g.Name + " (" + g.Genre + ")"
}

type PCEngineGame struct {
    Game
    Ean13 string
}

func handler(w http.ResponseWriter, r *http.Request) {
    gradius := PCEngineGame{Game{"Gradius", "STG"}, "4 988602 585032"}
    fmt.Fprintf(w, gradius.ToString())
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
