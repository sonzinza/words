# Words
This project call to Occurrence service, define [here](https://github.com/sonzinza/occurrence), pass text from input file ‘GoLang_Test.txt’ and prints JSON output returned from the service

# Test

Run this service first
```
    go mod tidy
    go run .
```

Test with curl or Postman
```
curl --location --request POST 'http://localhost:8080/words' \
--header 'Content-Type: application/json' \
--form 'txt_file=@"/path/to/words/GoLang_Test.txt"'
```

Result in json

```json
[
    {
        "word": "sed",
        "occurrence": 249
    },
    {
        "word": "in",
        "occurrence": 229
    },
    {
        "word": "ut",
        "occurrence": 214
    },
    {
        "word": "et",
        "occurrence": 204
    },
    {
        "word": "ac",
        "occurrence": 201
    },
    {
        "word": "non",
        "occurrence": 197
    },
    {
        "word": "eget",
        "occurrence": 191
    },
    {
        "word": "quis",
        "occurrence": 171
    },
    {
        "word": "id",
        "occurrence": 170
    },
    {
        "word": "amet",
        "occurrence": 169
    }
]
```