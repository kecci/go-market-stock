package gomarketstock

import (
	"io/ioutil"
	"net/http"
)

func GetDataFromURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode == 404 {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
