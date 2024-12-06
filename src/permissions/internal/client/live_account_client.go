package client

import (
	"bytes"
	"concierge/permissions/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// LiveAccountClient is a struct for managing HTTP requests related to accounts.
// It is used to interact with the live account service API.
// Args:
//     None
// Returns:
//     None
type LiveAccountClient struct {
	BaseURL string // BaseURL is the root URL for the API
}

// NewLiveAccountClient initializes a new LiveAccountClient with a given base URL.
// Args:
//     baseURL (string): The base URL for the account service API.
// Returns:
//     *LiveAccountClient: A pointer to a new instance of LiveAccountClient.
func NewLiveAccountClient(baseURL string) *LiveAccountClient {
	return &LiveAccountClient{
		BaseURL: baseURL,
	}
}

// Get sends a GET request to the specified path on the live account service API and returns the response body.
// Args:
//     path (string): The relative path to append to the base URL for the GET request.
// Returns:
//     ([]byte, error): The response body as a byte slice and any error encountered.
func (cl *LiveAccountClient) Get(path string) ([]byte, error) {

	fullURL := fmt.Sprintf("%s%s", cl.BaseURL, path)
	resp, err := http.Get(fullURL)
	e := "No ErrOr"
	if err != nil {
		e = err.Error()
	}
	fmt.Printf("%s/%s :: %s\n", cl.BaseURL, path, e)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body) // Read the response body
	if err != nil {
		return nil, fmt.Errorf("could not read response")
	}

	defer resp.Body.Close() // Close the response body when done
	return body, nil
}

// Post sends a POST request to the specified path on the live account service API with the given body data.
// Args:
//     path (string): The relative path to append to the base URL for the POST request.
//     body (any): The request body to send in the POST request.
// Returns:
//     ([]byte, error): The response body as a byte slice and any error encountered.
func (cl *LiveAccountClient) Post(path string, body any) ([]byte, error) {

	fullURL := fmt.Sprintf("%s%s", cl.BaseURL, path)

	jsonData, err := json.Marshal(body) // Marshal the body into JSON
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(fullURL, "application/json", bytes.NewBuffer(jsonData)) // Send POST request
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body) // Read the response body
	if err != nil {
		return nil, fmt.Errorf("could not read response")
	}

	defer resp.Body.Close() // Close the response body when done
	return respBody, nil
}

// PostLoginAttempt sends a POST request to attempt a login with the provided credentials
// and returns the account information if successful.
// Args:
//     request (models.LoginAttempt): The login attempt request data.
// Returns:
//     (*models.Account, error): A pointer to the account object and any error encountered.
func (cl *LiveAccountClient) PostLoginAttempt(request models.LoginAttempt) (*models.Account, error) {

	var loginResponse models.AccountResponse
	respBody, err := cl.Post("/accounts/login_attempt", request)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &loginResponse) // Parse the response body into the Account struct
	if err != nil {
		return nil, fmt.Errorf("could not parse response from json")
	}

	if loginResponse.AccountData == nil {
		return nil, fmt.Errorf("login failed")
	}

	fmt.Println(loginResponse.AccountData)

	return loginResponse.AccountData, nil
}
