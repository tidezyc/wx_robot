// Package main provides ...
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var cookies map[string]*http.Cookie

func httpGetBytes(url string) []byte {
	log.Println("http get:" + url)
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}
	req, err := http.NewRequest("get", fmt.Sprintf("%s", url), nil)
	if err != nil {
		log.Fatalln(err)
	}

	return execHttpReq(req)
}

func httpGetString(url string) string {
	res := httpGetBytes(url)
	if res != nil {
		return string(res)
	}
	return ""
}

func httpPostBytes(url string, params map[string]interface{}) []byte {
	log.Println("http post:" + url)
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}
	data, err := json.Marshal(params)
	if err != nil {
		log.Println(err)
		return nil
	}

	log.Println("post data: " + string(data))
	req, err := http.NewRequest("post", fmt.Sprintf("%s", url), bytes.NewBuffer(data))
	if err != nil {
		log.Fatalln(err)
	}

	return execHttpReq(req)
}

func httpPostString(url string, params map[string]interface{}) string {
	data := httpPostBytes(url, params)
	if data != nil {
		return string(data)
	}
	return ""
}
func execHttpReq(req *http.Request) []byte {
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.49 Safari/537.36")
	for _, v := range cookies {
		req.AddCookie(v)
	}
	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	defer res.Body.Close()
	for _, c := range res.Cookies() {
		if cookies == nil {
			cookies = make(map[string]*http.Cookie, 10)
		}
		cookies[c.Name] = c
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return data
}
