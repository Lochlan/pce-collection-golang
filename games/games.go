package games

type Game struct {
    Name string
}

func (g *Game) ToString() string {
    return g.Name
}

type PCEngineGame struct {
    Game
    Ean13 string
}
