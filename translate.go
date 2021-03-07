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

var filename = flag.String("file", "", "filename")
var text = flag.String("text", "", "text")
var tl = flag.String("tl", "", "Target language")
var sl = flag.String("sl", "", "Source language")

func main() {
	flag.Parse()
	content := *new(Content)

	if *text == "" {
		data, err := ioutil.ReadFile(*filename)
		check(err)
		*text = string(data)
	}

	baseUrl := "https://translate.googleapis.com/translate_a/single?client=gtx"
	query := url.QueryEscape(*text)
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
	fmt.Println("################### Original text #########################")
	fmt.Println(sourceText)
	fmt.Println("################### Target text #########################")
	fmt.Println(targetText)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
