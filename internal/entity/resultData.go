package entity

type ResultData struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	FinalWord    string         `json:"finalWord"`
	Columns      []ResultColumn `json:"columns" gorm:"serializer:json"`
	PuzzleDataID uint           `json:"puzzleDataId" gorm:"not null;"`
}

type ResultColumn struct {
	Position  int    `json:"position"`
	FinalWord string `json:"finalWord"`
}
