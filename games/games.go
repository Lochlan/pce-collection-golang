package games

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
