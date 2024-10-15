package client

import (
	"bytes"
	"concierge/permissions/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LiveAccountClient struct {
	BaseURL string
}

func NewLiveAccountClient(baseURL string) *LiveAccountClient {
	return &LiveAccountClient{
		BaseURL: baseURL,
	}
}

func (cl *LiveAccountClient) Get(path string) ([]byte, error) {
	fullURL := fmt.Sprintf("%s/%s", cl.BaseURL, path)
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not read response")
	}

	defer resp.Body.Close()
	return body, nil
}

func (cl *LiveAccountClient) Post(path string, body any) ([]byte, error) {
	fullURL := fmt.Sprintf("%s/%s", cl.BaseURL, path)

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(fullURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not read response")
	}

	defer resp.Body.Close()
	return respBody, nil
}

func (cl *LiveAccountClient) PostLoginAttempt(request models.LoginAttempt) (*models.Account, error) {
	var loginResponse models.Account
	respBody, err := cl.Post("/login-attempt", request)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &loginResponse)
	if err != nil {
		return nil, fmt.Errorf("Could not parse response from json")
	}

	return &loginResponse, nil
}
