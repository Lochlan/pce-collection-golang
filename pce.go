package main

import (
    "fmt"
    "net/http"
    "./games"
)

func handler(w http.ResponseWriter, r *http.Request) {
    gradius := &games.Game{
        Name: "Gradius",
    }
    fmt.Fprintf(w, gradius.ToString())
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
