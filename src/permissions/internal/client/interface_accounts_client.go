package client

import "concierge/permissions/internal/models"

type AccountClient interface {
	Get(string) ([]byte, error)
	Post(string, interface{}) ([]byte, error)
	PostLoginAttempt(models.LoginAttempt) (*models.Account, error)
}
