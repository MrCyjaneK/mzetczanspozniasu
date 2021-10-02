package mzklib

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func httpget(url string, dst interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}
