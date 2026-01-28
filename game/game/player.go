package game

type Player struct {
	Name     string   `json:"name"`
	Location Location `json:"location"`
}
type Location struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p *Player) Serialize() PlayerSaveData {
	return PlayerSaveData{
		Name:     p.Name,
		Location: p.Location,
	}
}
func DeserializePlayer(data PlayerSaveData) Player {
	return Player{
		Name:     data.Name,
		Location: data.Location,
	}
}
