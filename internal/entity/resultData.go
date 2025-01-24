package entity

import (
	"github.com/google/uuid"
	"github.com/kamva/mgm/v3"
)

type ResultData struct {
	mgm.DefaultModel `bson:",inline"`
	ID               uuid.UUID      `json:"id" bson:"id"`
	FinalWord        string         `json:"finalWord" bson:"finalWord"`
	Columns          []ResultColumn `json:"columns" bson:"columns"`
	PuzzleDataID     uuid.UUID      `json:"puzzleDataId" bson:"puzzleDataId"`
	UnknownWords     []string       `json:"unknownWords" bson:"unknownWords"`
}

type ResultColumn struct {
	Position  int    `json:"position" bson:"position"`
	FinalWord string `json:"finalWord" bson:"finalWord"`
}
