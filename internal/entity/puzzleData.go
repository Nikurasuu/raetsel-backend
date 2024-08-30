package entity

type PuzzleData struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	BridgeWords []string       `json:"bridgeWords"  gorm:"serializer:json"`
	Columns     []PuzzleColumn `json:"columns"  gorm:"serializer:json"`
}

type PuzzleColumn struct {
	Position        int    `json:"position"`
	First           string `json:"first"`
	Second          string `json:"second"`
	Space           int    `json:"space"`
	WantedCharacter int    `json:"wantedCharacter"`
}
