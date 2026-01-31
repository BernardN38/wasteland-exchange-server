package game

import "sync"

type Player struct {
	Name     string   `json:"name"`
	Location Location `json:"location"`
	sync.RWMutex
}
type Location struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p *Player) Serialize() PlayerSaveData {
	p.RLock()
	defer p.RUnlock()
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
