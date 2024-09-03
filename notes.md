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