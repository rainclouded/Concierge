package client

import (
	"concierge/permissions/internal/models"
	"fmt"
)

type MockAccountClient struct {
	accounts []*models.Account
}

func NewMockAccountClient() *MockAccountClient {
	return &MockAccountClient{
		accounts: []*models.Account{
			{ID: 1, Name: "admin"},
			{ID: 2, Name: "editor"},
			{ID: 3, Name: "viewer"}},
	}
}

func (cl *MockAccountClient) PostLoginAttempt(loginRequest models.LoginAttempt) (*models.Account, error) {
	if loginRequest.Username == loginRequest.Password {
		for _, acc := range cl.accounts {
			if acc.Name == loginRequest.Username {
				return acc, nil
			}
		}
	}

	return nil, fmt.Errorf("Invalid username or password")
}

func (cl *MockAccountClient) Get(path string) ([]byte, error) {
	return nil, fmt.Errorf("Cannot make arbitrary GET requests to Mock clients")
}

func (cl *MockAccountClient) Post(path string, body interface{}) ([]byte, error) {
	return nil, fmt.Errorf("Cannot make arbitrary GET requests to Mock clients")
}
