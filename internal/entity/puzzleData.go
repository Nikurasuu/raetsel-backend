package entity

import "github.com/kamva/mgm/v3"

type PuzzleData struct {
	mgm.DefaultModel `bson:",inline"`
	ID               uint           `json:"id" bson:"id"`
	BridgeWords      []string       `json:"bridgeWords"  bson:"bridgeWords"`
	Columns          []PuzzleColumn `json:"columns"  bson:"columns"`
}

type PuzzleColumn struct {
	mgm.DefaultModel `bson:",inline"`
	Position         int    `json:"position" bson:"position"`
	First            string `json:"first" bson:"first"`
	Second           string `json:"second" bson:"second"`
	Space            int    `json:"space" bson:"space"`
	WantedCharacter  int    `json:"wantedCharacter" bson:"wantedCharacter"`
}
