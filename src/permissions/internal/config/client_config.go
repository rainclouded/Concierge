package config

import (
	"concierge/permissions/internal/client"
	"fmt"
	"os"
)

// LoadAccountEndpoint loads the account endpoint client based on environment variables.
// It attempts to connect to a live account client if a valid endpoint is provided.
// Args:
//     None
// Returns:
//     client.AccountClient: The live account client if the endpoint is valid, otherwise a mock account client.
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

// TestAccountEndpoint tests if the given account client can successfully make a GET request to the "/healthcheck" endpoint.
// Args:
//     client (client.AccountClient): The account client to test.
// Returns:
//     bool: True if the endpoint is reachable and responds without error, otherwise false.
func TestAccountEndpoint(client client.AccountClient) bool {

	_, err := client.Get("/accounts")
	return err == nil

}
