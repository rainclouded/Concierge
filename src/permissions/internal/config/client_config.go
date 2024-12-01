package config

import (
	"concierge/permissions/internal/client"
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
	accEndpoint := os.Getenv("ACCOUNT_ENDPOINT") // Get the account endpoint from environment variables
	accCli = client.NewLiveAccountClient(accEndpoint) // Initialize the live account client
	if accEndpoint != "" && TestAccountEndpoint(accCli) { // Check if the endpoint is valid and reachable
		return accCli // Return the live account client if it's valid
	}

	return client.NewMockAccountClient() // Return the mock account client if the endpoint is not valid or not provided
}

// TestAccountEndpoint tests if the given account client can successfully make a GET request to the "/healthcheck" endpoint.
// Args:
//     client (client.AccountClient): The account client to test.
// Returns:
//     bool: True if the endpoint is reachable and responds without error, otherwise false.
func TestAccountEndpoint(client client.AccountClient) bool {
	_, err := client.Get("/healthcheck") // Attempt to get healthcheck from the account endpoint
	return err == nil // Return true if no error occurred, otherwise false
}
