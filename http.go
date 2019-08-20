package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	postCT  = "application/json;charset=UTF-8"
	timeout = 5 * time.Second
)

func get() (string, error) {
	req, err := http.NewRequest("GET", *thirdPart, nil)
	if err != nil {
		return "", err
	}
	client := &http.Client{Timeout: timeout}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	c, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(c), nil
}

func post(msg map[string]string) (map[string]string, error) {
	r, err := newReader(msg)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", *server, r)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var res map[string]string
	json.NewDecoder(resp.Body).Decode(&res)
	return res, nil
}

func newReader(msg interface{}) (io.Reader, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(data), nil
}
