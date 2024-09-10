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

func (wl *WordList) WordExists(word string) (bool, error) {
	rows, err := wl.db.Query("SELECT wort FROM words WHERE `wort` = '" + word + "'")
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return rows.Next(), nil
}

func (wl *WordList) SearchWord(word string) ([]string, error) {
	rows, err := wl.db.Query("SELECT wort FROM words WHERE `wort` LIKE '%" + word + "%'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []string
	for rows.Next() {
		var result string
		if err := rows.Scan(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func (wl *WordList) LeftWordWithBridgeWordExist(leftWord, bridgeWord string) (bool, error) {
	rows, err := wl.db.Query("SELECT wort FROM words WHERE `wort` = '" + leftWord + bridgeWord + "'")
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return rows.Next(), nil
}

func (wl *WordList) RightWordWithBridgeWordExist(rightWord, bridgeWord string) (bool, error) {
	rows, err := wl.db.Query("SELECT wort FROM words WHERE `wort` = '" + bridgeWord + rightWord + "'")
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return rows.Next(), nil
}
