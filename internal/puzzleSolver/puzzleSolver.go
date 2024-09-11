package puzzlesolver

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nikurasuu/raetsel-backend/internal/entity"
	"github.com/nikurasuu/raetsel-backend/internal/wordlist"
	"github.com/sirupsen/logrus"
)

type puzzleSolver struct {
	logger   *logrus.Logger
	wordlist *wordlist.WordList
}

func NewPuzzleSolver(logger *logrus.Logger, wordlist *wordlist.WordList) *puzzleSolver {
	return &puzzleSolver{
		logger:   logger,
		wordlist: wordlist,
	}
}

func (p *puzzleSolver) solveColumn(leftWord string, rightWord string, wantedCharacters int, bridgeWords []string) (string, error) {
	startTime := time.Now()

	for _, bridgeWord := range bridgeWords {
		if len(bridgeWord) == wantedCharacters {
			if p.wordlist.LeftWordWithBridgeWordExist(leftWord, bridgeWord) && p.wordlist.RightWordWithBridgeWordExist(rightWord, bridgeWord) {
				elapsedTime := time.Since(startTime)
				p.logger.Debug("solveColumn found ", elapsedTime)
				return bridgeWord, nil
			}
		}
	}

	elapsedTime := time.Since(startTime)
	p.logger.Debug("solveColumn not found ", elapsedTime)
	return "", fmt.Errorf("no valid bridge word found for %s and %s", leftWord, rightWord)
}

func (p *puzzleSolver) SolvePuzzle(puzzle entity.PuzzleData) (entity.ResultData, error) {
	startTime := time.Now()

	var resultData entity.ResultData
	resultData.ID = uuid.New()
	resultData.PuzzleDataID = puzzle.ID

	for _, column := range puzzle.Columns {
		bridgeWord, err := p.solveColumn(column.First, column.Second, column.Space, puzzle.BridgeWords)
		if err != nil {
			p.logger.Error(err)
		}
		resultColumn := entity.ResultColumn{
			Position:  column.Position,
			FinalWord: bridgeWord,
		}
		resultData.Columns = append(resultData.Columns, resultColumn)
	}

	elapsedTime := time.Since(startTime)
	p.logger.Debug("SolvePuzzle took ", elapsedTime)
	return resultData, nil
}
