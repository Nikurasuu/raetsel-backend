package entity

type PuzzleData struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	BridgeWords []string `json:"bridgeWords"`
	Columns     []Column `json:"columns"`
}

type Column struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	Position        int    `json:"position"`
	First           string `json:"first"`
	Second          string `json:"second"`
	Space           int    `json:"space"`
	WantedCharacter int    `json:"wantedCharacter"`
}
