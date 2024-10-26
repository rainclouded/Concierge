package config

import (
	"concierge/permissions/internal/client"
	"os"
)

func LoadAccountEndpoint() client.AccountClient {
	var accCli client.AccountClient
	accEndpoint := os.Getenv("ACCOUNT_ENDPOINT")
	accCli = client.NewLiveAccountClient(accEndpoint)
	if accEndpoint != "" && TestAccountEndpoint(accCli) {
		return accCli
	}

	return client.NewMockAccountClient()
}

func TestAccountEndpoint(client client.AccountClient) bool {
	_, err := client.Get("/healthcheck")
	return err == nil
}
