package game

type Settlement struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Location   Location `json:"location"`
	Population int      `json:"population"`
}

func NewSettlement(name string, x, y int) *Settlement {
	return &Settlement{
		Name:       name,
		Location:   Location{X: x, Y: y},
		Population: 0,
	}
}

func (s *Settlement) Tick() {
	s.Population += 1
}

func (s *Settlement) Serialize() SettlementSaveData {
	return SettlementSaveData{
		Population: s.Population,
	}
}

func DeserializeSettlement(data SettlementSaveData) Settlement {
	return Settlement{
		Population: data.Population,
	}
}
