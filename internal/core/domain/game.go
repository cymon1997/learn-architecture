package domain

type Game struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	State         string        `json:"state"`
	BoardSettings BoardSettings `json:"board_settings"`
	Board         Board         `json:"board"`
}

func NewGame(name string, size, bombs uint) Game {
	// TODO: Implement initial game setup
	return Game{
		ID:    "",
		Name:  name,
		State: "",
		BoardSettings: BoardSettings{
			Size:  size,
			Bombs: bombs,
		},
		Board: nil,
	}
}
