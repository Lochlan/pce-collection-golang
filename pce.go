package main

import (
    "fmt"
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

func main() {
    gradius := PCEngineGame{Game{"Gradius", "STG"}, "4 988602 585032"}
    fmt.Println(gradius.ToString())
}
