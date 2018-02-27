package mlib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpPost(url string, reqs, resp interface{}, timeout int) error {
	body, err := json.Marshal(reqs)
	if err != nil {
		return err
	}
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	raw, err := client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer raw.Body.Close()

	body, err = ioutil.ReadAll(raw.Body)
	if err != nil {
		return err
	}

	if len(body) != 0 && resp != nil {
		err = json.Unmarshal(body, resp)
		if err != nil {
			return err
		}
	}

	return nil
}

func HttpDelete(url string, resp interface{}, timeout int) error {
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	raw, err := client.Do(request)
	if err != nil {
		return err
	}
	defer raw.Body.Close()

	body, err := ioutil.ReadAll(raw.Body)
	if err != nil {
		return err
	}

	if len(body) != 0 && resp != nil {
		err = json.Unmarshal(body, resp)
		if err != nil {
			return err
		}
	}

	return nil
}
