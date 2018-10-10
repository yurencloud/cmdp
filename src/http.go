package src

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

func Http(method string, path string, param string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest(method, URL+path, strings.NewReader(param))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	token, _ := ReadToken() //从文件取出数据
	req.Header.Set("Authorization", "Bearer "+string(token))
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}

func HttpJson(method string, path string, param []byte) []byte {
	client := &http.Client{}
	req, err := http.NewRequest(method, URL+path, bytes.NewBuffer(param))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	token, _ := ReadToken() //从文件取出数据
	req.Header.Set("Authorization", "Bearer "+string(token))
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}
