package puzzlesolver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/nikurasuu/raetsel-backend/internal/entity"
	"github.com/nikurasuu/raetsel-backend/internal/wordlist"
	"github.com/sirupsen/logrus"
)

type puzzleSolver struct {
	logger     *logrus.Logger
	wordlist   *wordlist.WordList
	httpClient *http.Client
}

func NewPuzzleSolver(logger *logrus.Logger, wordlist *wordlist.WordList) *puzzleSolver {
	httpClient := &http.Client{}
	return &puzzleSolver{
		logger:     logger,
		wordlist:   wordlist,
		httpClient: httpClient,
	}
}

type apiResponse struct {
	Left   string `json:"left"`
	Bridge string `json:"bridge"`
	Right  string `json:"right"`
}

func (p *puzzleSolver) solveColumn(leftWord string, rightWord string, wantedCharacters int, bridgeWords []string) (string, error) {
	for _, bridgeWord := range bridgeWords {
		if len(bridgeWord) == wantedCharacters {
			if p.wordlist.LeftWordWithBridgeWordExist(leftWord, bridgeWord) && p.wordlist.RightWordWithBridgeWordExist(rightWord, bridgeWord) {
				return bridgeWord, nil
			}
		}
	}

	bridgeWord, err := p.solveColumnWithAPI(leftWord, rightWord, wantedCharacters)
	p.logger.Infof("Could not solve column with wordlist, trying API: %s %s", leftWord, rightWord)
	if err != nil {
		return "", fmt.Errorf("error solving column with API: %v", err)
	} else if bridgeWord != "" {
		return bridgeWord, nil
	}

	return "", fmt.Errorf("no valid bridge word found for %s and %s", leftWord, rightWord)
}

func (p *puzzleSolver) solveColumnWithAPI(leftWord, rightWord string, wantedCharacters int) (string, error) {
	url := fmt.Sprintf("https://api.kwdb.ch/api/bridge-builder/?left=%s&right=%s&chars=%d", leftWord, rightWord, wantedCharacters)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request to bridge word API: %v", err)
	}

	res, err := p.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error calling bridge word API: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	var apiResponses []apiResponse
	if err := json.Unmarshal(data, &apiResponses); err != nil {
		return "", fmt.Errorf("error unmarshalling response body: %v", err)
	}

	if len(apiResponses) == 0 {
		return "", fmt.Errorf("no valid bridge word found for %s and %s", leftWord, rightWord)
	}

	// lowercase the bridge word
	bridgeWord := strings.ToLower(apiResponses[0].Bridge)

	return bridgeWord, nil
}

func (p *puzzleSolver) SolvePuzzle(puzzle *entity.PuzzleData) (entity.ResultData, error) {
	startTime := time.Now()
	p.logger.Info("Solving puzzle with ID: ", puzzle.ID)

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
	p.logger.Info("Solved puzzle with ID: ", puzzle.ID)
	return resultData, nil
}
