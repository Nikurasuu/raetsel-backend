package entity

import (
	"github.com/google/uuid"
	"github.com/kamva/mgm/v3"
)

type PuzzleData struct {
	mgm.DefaultModel `bson:",inline"`
	ID               uuid.UUID      `json:"id" bson:"id"`
	BridgeWords      []string       `json:"bridgeWords"  bson:"bridgeWords"`
	Columns          []PuzzleColumn `json:"columns"  bson:"columns"`
	// ResultDataId     uuid.UUID      `json:"resultDataId" bson:"resultDataId"`
}

type PuzzleColumn struct {
	Position        int    `json:"position" bson:"position"`
	First           string `json:"first" bson:"first"`
	Second          string `json:"second" bson:"second"`
	Space           int    `json:"space" bson:"space"`
	WantedCharacter int    `json:"wantedCharacter" bson:"wantedCharacter"`
}
