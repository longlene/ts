package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Trans struct {
	ErrorCode       int32                  `json:"errorCode"`
	ElapsedTime     int32                  `json:"elapsedTime"`
	Type            string                 `json:"type"`
	TranslateResult interface{}            `json:"translateResult"`
	SmartResult     map[string]interface{} `json:"smartResult"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ts word")
		os.Exit(1)
	}

	query(os.Args[1])
}

func query(word string) {
	const URL string = "http://fanyi.youdao.com/translate"

	ext := url.Values{}
	ext.Add("smartresult", "dict")
	ext.Add("smartresult", "rule")
	ext.Add("smartresult", "ugc")
	ext.Set("sessionFrom", "dict.top")

	data := url.Values{}
	data.Set("type", "AUTO")
	data.Set("i", word)
	data.Set("doctype", "json")
	data.Set("xmlVersion", "1.4")
	data.Set("keyfrom", "fanyi.web")
	data.Set("ue", "UTF-8")
	data.Set("typoResult", "true")
	data.Set("flag", "false")

	var ur string = URL + "?" + ext.Encode()
	resp, err := http.PostForm(ur, data)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var tr Trans
	err = json.Unmarshal(body, &tr)
	if err != nil {
		panic(err)
	}

	v := tr.SmartResult["entries"]
	if v != nil {
		fmt.Printf("[%s]\n", word)

		entries := v.([]interface{})
		for v := range entries {
			s := entries[v].(string)
			if len(s) > 0 {
				fmt.Printf("\033[1;32m%s\033[0m\n", s)
			}
		}
	}
}
