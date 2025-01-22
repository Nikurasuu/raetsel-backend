package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/otiai10/gosseract/v2"
)

type CrosswordData struct {
	LeftWord       string
	RightWord      string
	Length         int
	MarkedPosition int
}

func main() {
	client := gosseract.NewClient()
	defer client.Close()

	imagePath := "test.png"

	client.SetImage(imagePath)

	text, err := client.Text()
	if err != nil {
		log.Fatalf("Failed to perform OCR: %v", err)
	}

	fmt.Println("Extracted Text:")
	fmt.Println(text)

	extractedData := parseCrosswordData(text)

	for _, data := range extractedData {
		fmt.Printf("Left: %s, Right: %s, Length: %d, Marked Position: %d\n",
			data.LeftWord, data.RightWord, data.Length, data.MarkedPosition)
	}

}

func parseCrosswordData(text string) []CrosswordData {
	lines := splitIntoLines(text)
	var data []CrosswordData

	for _, line := range lines {
		re := regexp.MustCompile(`^(\\w+)\\s+(\\s+)\\s+(\\w+)$`)
		matches := re.FindStringSubmatch(line)
		if matches == nil || len(matches) < 4 {
			continue
		}

		leftWord := matches[1]
		spaceBetween := matches[2]
		rightWord := matches[3]

		length := len(spaceBetween)

		markedPosition := strings.Index(spaceBetween, "^") + 1

		data = append(data, CrosswordData{
			LeftWord:       leftWord,
			RightWord:      rightWord,
			Length:         length,
			MarkedPosition: markedPosition,
		})
	}

	return data
}

func splitIntoLines(text string) []string {
	lines := strings.Split(text, "\n")
	var filteredLines []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			filteredLines = append(filteredLines, line)
		}
	}

	return filteredLines
}
