using Newtonsoft.Json;
using task_system_server.Models;

namespace task_system_server.Clients
{
    /* 
    PermissionClient is a class responsible for interacting with the Permissions service.
    It uses IHttpClientFactory to create HTTP clients for making API calls to the service.
    */
    public class PermissionClient
    {
        private readonly IHttpClientFactory _httpClientFactory;

        /* 
        Constructor for PermissionClient.
        Initializes the HTTP client factory for creating HTTP clients.
        
        Args:
        httpClientFactory: A factory for creating HTTP client instances.
        */
        public PermissionClient(IHttpClientFactory httpClientFactory)
        {
            _httpClientFactory = httpClientFactory;
        }

        /* 
        GetDataFromServiceAsync makes an asynchronous GET request to the given endpoint with the provided API key.
        It logs the connection attempt and response from the service.

        Args:
        endpoint: The endpoint to send the GET request to.
        apiKey: The API key for authenticating the request.
        
        Returns:
        A string containing the response from the service.
        */
        public async Task<string> GetDataFromServiceAsync(string endpoint, string apiKey)
        {
            Console.WriteLine($"Connecting to Permissions @ {endpoint}");
            var client = _httpClientFactory.CreateClient();
            client.DefaultRequestHeaders.Add("X-API-Key", apiKey);
            var response = await client.GetAsync(endpoint);

            var resp = await response.Content.ReadAsStringAsync();
            Console.WriteLine("Successfully connected");
            return resp;
        }

        /* 
        GetSessionData retrieves session data for the user associated with the provided session key.
        It makes a request to the session endpoint and returns the session data if successful.

        Args:
        sessionKey: The session key used for authenticating the request.
        
        Returns:
        A SessionData object containing the session information, or null if the deserialization fails.
        */
        public async Task<SessionData?> GetSessionData(string sessionKey)
        {
            var endpoint = Environment.GetEnvironmentVariable("SESSIONS_ENDPOINT");
            var response = await GetDataFromServiceAsync(endpoint + "/sessions/me", sessionKey);
            Console.WriteLine($"Response: {response}");
            return JsonConvert.DeserializeObject<SessionData>(response);
        }
    }
}
