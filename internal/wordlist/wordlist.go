package wordlist

import (
	"database/sql"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type WordList struct {
	db *sql.DB
}

func NewWordList(dbPath string) (*WordList, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return &WordList{db: db}, nil
}

// This is needed because the database only contains words with umlauts
// but the puzzle data might not contain umlauts.
// This ensures that the search and works with umlauts.
func replaceUmlauts(word string) string {
	replacements := map[string]string{
		"UE": "Ü",
		"AE": "Ä",
		"OE": "Ö",
		"ue": "ü",
		"ae": "ä",
		"oe": "ö",
	}

	for old, new := range replacements {
		word = strings.ReplaceAll(word, old, new)
	}

	return word
}

func (wl *WordList) WordExists(word string) bool {
	query := "SELECT wort FROM words WHERE LOWER(wort) = LOWER(?)"
	word = replaceUmlauts(word)
	rows, err := wl.db.Query(query, word)
	if err != nil {
		return false
	}
	defer rows.Close()

	return rows.Next()
}

func (wl *WordList) SearchWord(word string) []string {
	query := "SELECT wort FROM words WHERE LOWER(wort) LIKE '%' || LOWER(?) || '%'"
	word = replaceUmlauts(word)
	rows, err := wl.db.Query(query, word)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var results []string
	for rows.Next() {
		var result string
		if err := rows.Scan(&result); err != nil {
			return nil
		}
		results = append(results, result)
	}
	return results
}

func (wl *WordList) LeftWordWithBridgeWordExist(leftWord, bridgeWord string) bool {
	query := "SELECT wort FROM words WHERE LOWER(wort) = LOWER(?)"
	leftWord = replaceUmlauts(leftWord)
	bridgeWord = replaceUmlauts(bridgeWord)
	rows, err := wl.db.Query(query, leftWord+bridgeWord)
	if err != nil {
		return false
	}
	defer rows.Close()

	return rows.Next()
}

func (wl *WordList) RightWordWithBridgeWordExist(rightWord, bridgeWord string) bool {
	query := "SELECT wort FROM words WHERE LOWER(wort) = LOWER(?)"
	rightWord = replaceUmlauts(rightWord)
	bridgeWord = replaceUmlauts(bridgeWord)
	rows, err := wl.db.Query(query, bridgeWord+rightWord)
	if err != nil {
		return false
	}
	defer rows.Close()

	return rows.Next()
}
