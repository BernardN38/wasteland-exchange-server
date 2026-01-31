package service

import "github.com/BernardN38/wasteland-exchange-server/game"

type Service struct {
	Game *game.Game
}

func NewService(game *game.Game) *Service {
	return &Service{
		Game: game,
	}
}

func (s *Service) GetHealth() (string, error) {
	return "Service is healthy", nil
}

func (s *Service) StartGame() (game.StartGameResponse, error) {
	// Logic to start the game would go here
	newGame, err := game.LoadGame()
	if err != nil {
		return game.StartGameResponse{Players: []game.Player{}}, err
	}
	s.Game = newGame
	return game.StartGameResponse{Players: newGame.Players, Settlements: newGame.Settlements}, nil
}
