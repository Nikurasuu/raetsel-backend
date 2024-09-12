1. Ansatz für die Lösung des Problems

```
def searchColumn(leftWord, rightWord, wantedCharacters, bridgeWords):
    for bridgeWord in bridgeWords:
        if len(bridgeWord) == wantedCharacters:
            search for leftWord+bridgeWord in Wordlist:
                if exists:
                    search for rightWord+bridgeWord in Wordlist:
                        if exists:
                            return bridgeWord
                        if not exists:
                            do nothing
                if not exists:
                    do nothing
    return None
```

2. Ablauf des Programms
    1. Bild kommt an
    2. Bild wird in PuzzleData Objekt umgewandelt
    3. PuzzleData Objekt wird in die Datenbank eingetragen (puzzleDataHandler.CreatePuzzleData) 
    4. ResultData Objekt wird aus PuzzleData Objekt erstellt (puzzleSolver.SolvePuzzle)
    5. ResultData Objekt wird in die Datenbank eingetragen (resultDataHandler.CreateResultData)