package entity

import "github.com/kamva/mgm/v3"

type ResultData struct {
	mgm.DefaultModel `bson:",inline"`
	ID               uint           `json:"id" bson:"id"`
	FinalWord        string         `json:"finalWord" bson:"finalWord"`
	Columns          []ResultColumn `json:"columns" bson:"columns"`
	PuzzleDataID     uint           `json:"puzzleDataId" bson:"puzzleDataId"`
}

type ResultColumn struct {
	mgm.DefaultModel `bson:",inline"`
	Position         int    `json:"position" bson:"position"`
	FinalWord        string `json:"finalWord" bson:"finalWord"`
}
