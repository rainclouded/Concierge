package config

import (
	"concierge/permissions/internal/client"
	"fmt"
	"os"
)

func LoadAccountEndpoint() client.AccountClient {
	var accCli client.AccountClient
	accEndpoint := os.Getenv("ACCOUNT_ENDPOINT")
	accCli = client.NewLiveAccountClient(accEndpoint)
	if accEndpoint != "" {
		fmt.Println("Connected to Account")
		return accCli
	} else {
		fmt.Println("Connected to Mock account")
		return client.NewMockAccountClient()
	}

}

func TestAccountEndpoint(client client.AccountClient) bool {
	_, err := client.Get("/accounts")
	return err == nil
}
