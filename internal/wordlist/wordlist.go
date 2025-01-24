package wordlist

import (
	"database/sql"

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

func (wl *WordList) WordExists(word string) bool {
	query := "SELECT wort FROM words WHERE LOWER(wort) = LOWER(?)"
	rows, err := wl.db.Query(query, word)
	if err != nil {
		return false
	}
	defer rows.Close()

	return rows.Next()
}

func (wl *WordList) SearchWord(word string) []string {
	query := "SELECT wort FROM words WHERE LOWER(wort) LIKE '%' || LOWER(?) || '%'"
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
	rows, err := wl.db.Query(query, leftWord+bridgeWord)
	if err != nil {
		return false
	}
	defer rows.Close()

	return rows.Next()
}

func (wl *WordList) RightWordWithBridgeWordExist(rightWord, bridgeWord string) bool {
	query := "SELECT wort FROM words WHERE LOWER(wort) = LOWER(?)"
	rows, err := wl.db.Query(query, bridgeWord+rightWord)
	if err != nil {
		return false
	}
	defer rows.Close()

	return rows.Next()
}
