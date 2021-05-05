package main

import (
	"net/http"
)

func GetRequest(url string) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
	return nil, err
	}
	req.Header.Add("X-Riot-Token", "RGAPI-4ba4f38c-5adc-4cf2-97db-d074cbc002e1")
	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return resp, nil
}
