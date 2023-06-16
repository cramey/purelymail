package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GenericResponse struct {
	Type    string `json:"type"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

const (
	MIME_JSON string = "application/json"
)

func (api *API) exec(ur, rm, ct string, bd interface{}) ([]byte, error) {
	var req *http.Request
	var err error
	if bd != nil {
		js, err := json.Marshal(bd)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(rm, ur, bytes.NewReader(js))
	} else {
		req, err = http.NewRequest(rm, ur, nil)
	}
	if err != nil {
		return nil, err
	}

	req.Header.Set("Purelymail-Api-Token", api.Token)
	req.Header.Set("Accept", ct)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		// Fallback error handling
		return body, fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}
