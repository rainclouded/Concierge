using amenities_server.model;
using Newtonsoft.Json;

namespace amenities_server.services
{
    public class PermissionClient
    {
        private readonly IHttpClientFactory _httpClientFactory;
        
        /*
        Constructor that initializes the PermissionClient with an IHttpClientFactory to create HTTP client instances.
        Args:
            httpClientFactory: An instance of IHttpClientFactory used to create HTTP client instances.
        */
        public PermissionClient(IHttpClientFactory httpClientFactory) 
        {
            _httpClientFactory = httpClientFactory;
        }

        /*
        Asynchronously retrieves data from a service endpoint with the provided API key for authorization.
        Args:
            endpoint: The endpoint URL to retrieve data from.
            apiKey: The API key used for authentication in the request headers.
        Returns:
            Task<string>: A task that represents the asynchronous operation, containing the response data as a string.
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
        Asynchronously retrieves session data for the given session key by calling the sessions endpoint.
        Args:
            sessionKey: The session key used to authenticate the request.
        Returns:
            Task<SessionData?>: A task that represents the asynchronous operation, containing the session data or null if not found.
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
