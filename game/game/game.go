package game

type Game struct {
	Players []Player // List of player names
}

func NewGame() (*Game, error) {
	return &Game{
		Players: []Player{{Name: "player1", Location: Location{X: 100, Y: 200}}, {Name: "player2", Location: Location{X: 150, Y: 250}}},
	}, nil
}
