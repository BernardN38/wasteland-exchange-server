package game

type StartGameResponse struct {
	Players     []Player      `json:"players"`
	Settlements []*Settlement `json:"settlements"`
}

type PlayerSaveData struct {
	Name     string   `json:"name"`
	Location Location `json:"location"`
}
type SettlementSaveData struct {
	Population int `json:"population"`
}
