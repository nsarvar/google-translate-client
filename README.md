# Client for Google Translate API

### Using

```
${PWD}

go build
./translate -text="Hello World" -sl=en -tl=uz
./translate -file=data.txt -sl=en -tl=uz

* text - text to translated
* file - filename with its path (if text is not provided, reads from file) 
* sl - source language
* tl - target language 
```
