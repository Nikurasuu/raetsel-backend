package puzzlesolver

import (
	"fmt"
	"time"

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

func (p *puzzleSolver) SolveColumn(leftWord string, rightWord string, wantedCharacters int, bridgeWords []string) (string, error) {
	startTime := time.Now()

	for _, bridgeWord := range bridgeWords {
		if len(bridgeWord) == wantedCharacters {
			if p.wordlist.LeftWordWithBridgeWordExist(leftWord, bridgeWord) && p.wordlist.RightWordWithBridgeWordExist(rightWord, bridgeWord) {
				elapsedTime := time.Since(startTime)
				fmt.Printf("Solving took %s\n", elapsedTime)
				return bridgeWord, nil
			}
		}
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("Solving took %s\n", elapsedTime)
	return "", fmt.Errorf("no valid bridge word found")
}
