package game

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Game struct {
	Players     []Player      // List of player names
	Settlements []*Settlement // List of settlements
	Ticker      time.Ticker   // Game loop ticker
}
type GameSave struct {
	Players     []PlayerSaveData              `json:"players"`
	Settlements map[string]SettlementSaveData `json:"settlements"`
}

func (game *Game) RunGameLoop() {
	for range game.Ticker.C {
		// Game loop logic would go here
		for _, settlement := range game.Settlements {
			settlement.Tick()
		}
		log.Println("tick")
		game.SaveGame()
	}
}

func NewGame() (*Game, error) {
	settlements, err := loadDefaultSettlements()
	if err != nil {
		return nil, err
	}
	return &Game{
		Players:     []Player{{Name: "player1", Location: Location{X: 100, Y: 200}}, {Name: "player2", Location: Location{X: 150, Y: 250}}},
		Settlements: settlements,
		Ticker:      *time.NewTicker(time.Millisecond * 100),
	}, nil
}

func LoadGame() (*Game, error) {
	saveData := GameSave{}
	// load save data from saves/game_save.json
	file, err := os.ReadFile("saves/save.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &saveData)
	if err != nil {
		return nil, err
	}

	settlementSaveData, err := loadSavedSettlements(&saveData)
	if err != nil {
		return nil, err
	}
	return &Game{
		Players:     []Player{{Name: "player1", Location: Location{X: 100, Y: 200}}, {Name: "player2", Location: Location{X: 150, Y: 250}}},
		Settlements: settlementSaveData,
	}, nil
}

func loadDefaultSettlements() ([]*Settlement, error) {
	var settlements []*Settlement
	file, err := os.ReadFile("common/settlements.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &settlements)
	if err != nil {
		return nil, err
	}
	return settlements, nil
}

func loadSavedSettlements(saveData *GameSave) ([]*Settlement, error) {
	var settlements []*Settlement
	// load default settlements from common/settlements.json
	defaultFile, err := os.ReadFile("common/settlements.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(defaultFile, &settlements)
	if err != nil {
		return nil, err
	}

	// apply save data
	for _, settlement := range settlements {
		savedData, exists := saveData.Settlements[settlement.ID]
		if exists {
			settlement.Population = savedData.Population
		}
	}
	return settlements, nil
}

func (game *Game) SaveGame() error {
	saveData := GameSave{
		Players:     []PlayerSaveData{},
		Settlements: make(map[string]SettlementSaveData),
	}
	for _, player := range game.Players {
		saveData.Players = append(saveData.Players, player.Serialize())
	}
	for _, settlement := range game.Settlements {
		saveData.Settlements[settlement.ID] = settlement.Serialize()
	}
	data, err := json.MarshalIndent(saveData, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("saves/save.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}
