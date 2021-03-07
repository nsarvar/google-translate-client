package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Content []interface{}

func main() {
	content := *new(Content)

	filename := flag.String("file", "data.txt", "filename")
	tl := flag.String("tl", "uz", "Target language")
	sl := flag.String("sl", "en", "Source language")

	data, err := ioutil.ReadFile(*filename)
	check(err)
	baseUrl := "https://translate.googleapis.com/translate_a/single?client=gtx"
	query := url.QueryEscape(string(data))
	translateUrl := fmt.Sprintf("%s&sl=%s&tl=%s&dt=t&q=%s", baseUrl, *sl, *tl, query)
	resp, err := http.Get(translateUrl)
	check(err)
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &content)
	check(err)
	targetText := ""
	sourceText := ""
	t := content[0].([]interface{})
	for x := range t {
		i := t[x].([]interface{})
		targetText += i[0].(string)
		sourceText += i[1].(string)
	}
	fmt.Printf("Original text: %s \n\n", sourceText)
	fmt.Printf("Translated text: %s \n", targetText)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
