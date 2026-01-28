package game

type Player struct {
	Name     string   `json:"name"`
	Location Location `json:"location"`
}
type Location struct {
	X int `json:"x"`
	Y int `json:"y"`
}
