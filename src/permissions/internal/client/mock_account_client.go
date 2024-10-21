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
	} else if loginRequest.Password == "" {
		if loginRequest.Username == "404" {
			return &models.Account{
				ID:   5,
				Name: "Guest #404",
			}, nil
		}
	}

	return nil, fmt.Errorf("invalid username or password")
}

func (cl *MockAccountClient) Get(path string) ([]byte, error) {
	return nil, fmt.Errorf("cannot make arbitrary GET requests to Mock clients")
}

func (cl *MockAccountClient) Post(path string, body interface{}) ([]byte, error) {
	return nil, fmt.Errorf("cannot make arbitrary GET requests to Mock clients")
}
