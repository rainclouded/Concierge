package client

import (
	"concierge/permissions/internal/models"
	"fmt"
)

// MockAccountClient is a mock implementation of an account client for testing purposes.
// It simulates account management functionality for testing purposes.
// Args:
//     None
// Returns:
//     None
type MockAccountClient struct {
	accounts []*models.Account // List of mock accounts
}

// NewMockAccountClient initializes a new MockAccountClient with some predefined accounts.
// Args:
//     None
// Returns:
//     *MockAccountClient: A pointer to a new instance of MockAccountClient with mock accounts.
func NewMockAccountClient() *MockAccountClient {
	return &MockAccountClient{
		accounts: []*models.Account{
			{ID: 1, Name: "admin"},
			{ID: 2, Name: "editor"},
			{ID: 3, Name: "viewer"},
			{ID: 4, Name: "404"}},
	}
}

// PostLoginAttempt simulates the process of logging in by verifying the username and password.
// If the username and password match, it returns the corresponding account.
// If the password is empty and the username is "404", it returns a guest account.
// Args:
//     loginRequest (models.LoginAttempt): The login request containing username and password.
// Returns:
//     (*models.Account, error): The account matching the username and password or an error if login fails.
func (cl *MockAccountClient) PostLoginAttempt(loginRequest models.LoginAttempt) (*models.Account, error) {
	if loginRequest.Username == loginRequest.Password {
		for _, acc := range cl.accounts {
			if acc.Name == loginRequest.Username {
				return acc, nil // Return matching account
			}
		}
	} else if loginRequest.Password == "" {
		if loginRequest.Username == "404" {
			return &models.Account{
				ID:   5,
				Name: "Guest #404",
			}, nil // Return guest account for username "404"
		}
	}

	return nil, fmt.Errorf("invalid username or password") // Return error for invalid credentials
}

// Get simulates an HTTP GET request, but returns an error as it is not supported by the mock client.
// Args:
//     path (string): The path for the GET request (not used in mock).
// Returns:
//     ([]byte, error): Always returns an error stating GET is not supported.
func (cl *MockAccountClient) Get(path string) ([]byte, error) {
	return nil, fmt.Errorf("cannot make arbitrary GET requests to Mock clients") // Return error for unsupported operation
}

// Post simulates an HTTP POST request, but returns an error as it is not supported by the mock client.
// Args:
//     path (string): The path for the POST request (not used in mock).
//     body (interface{}): The body of the POST request (not used in mock).
// Returns:
//     ([]byte, error): Always returns an error stating POST is not supported.
func (cl *MockAccountClient) Post(path string, body interface{}) ([]byte, error) {
	return nil, fmt.Errorf("cannot make arbitrary GET requests to Mock clients") // Return error for unsupported operation
}
