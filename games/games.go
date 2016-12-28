package games

type Game struct {
    Name string
}

func (g *Game) ToString() string {
    return g.Name
}
