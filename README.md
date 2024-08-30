POST /puzzle
```json
{
    "id": 1,
	"metadata": {},
	"status": "created",
}
```

GET /result?id=some-id
```json
{
	"puzzle":
	{
	    "id": 1,
	    "metadata": {},
	    "bridgeWords": [
	        "word1",
	        "word2",
	        "wordn"
	    ],
	    "columns": []
	},
	"result":
	{
		"id": 1,
		"metadata": {},
		"finalWord": "word1",
		"columns": []
	}
}
```

puzzle column:
```json
{
	"id": 1,
	"position": 1,
	"first": "word1",
	"second": "word2",
	"space": 4,
	"wantedCharacter": 3
}
```

result column:
```json
{
	"position": 1,
	"finalWord": "word2"
}
```
