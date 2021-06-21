package common

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func HTTPGet(client *http.Client, url string, result interface{}) (int, error) {
	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}

	return resp.StatusCode, json.Unmarshal(body, result)
}
