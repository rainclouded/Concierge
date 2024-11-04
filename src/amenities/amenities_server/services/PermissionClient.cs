using amenities_server.model;
using Newtonsoft.Json;

namespace amenities_server.services
{
    public class PermissionClient
    {
        private readonly IHttpClientFactory _httpClientFactory;
        
        public PermissionClient(IHttpClientFactory httpClientFactory) 
        {
            _httpClientFactory = httpClientFactory;
        }

        public async Task<string> GetDataFromServiceAsync(string endpoint, string apiKey)
        {
            Console.WriteLine($"Connecting to Permissions @ {endpoint}");
            var client = _httpClientFactory.CreateClient();
            client.DefaultRequestHeaders.Add("X-Api-Key", apiKey);
            var response = await client.GetAsync(endpoint);

            var resp = await response.Content.ReadAsStringAsync();
            Console.WriteLine("Successfully connected");
            return resp;
        }

        public async Task<SessionData?> GetSessionData(string sessionKey)
        {
            var endpoint = Environment.GetEnvironmentVariable("SESSIONS_ENDPOINT");
            var response = await GetDataFromServiceAsync(endpoint + "/sessions/me", sessionKey);
            Console.WriteLine($"Response: {response}"); 
            return JsonConvert.DeserializeObject<SessionData>(response);
        }
    }
}
