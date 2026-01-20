package game

type Game struct {
	Players []string // List of player names
}

func NewGame() (*Game, error) {
	return &Game{
		Players: []string{"player1", "player2"},
	}, nil
}
