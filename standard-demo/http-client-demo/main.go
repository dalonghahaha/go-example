package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	httpGet()
	httpPost()
}

//httpget
func httpGet() {
	resp, err := http.Get("http://127.0.0.1:8000/get")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

//httppost
func httpPost() {
	resp, err := http.PostForm("http://127.0.0.1:8000/post", url.Values{"id": {"1"}, "name": {"hello"}})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
