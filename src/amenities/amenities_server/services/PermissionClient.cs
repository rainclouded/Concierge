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

        public async Task<string> GetDataFromServiceAsync(string endpoint)
        {
            Console.WriteLine($"Connecting to Permissions @ {endpoint}");
            var client = _httpClientFactory.CreateClient();
            var response = await client.GetAsync(endpoint);

            var resp = await response.Content.ReadAsStringAsync();
            Console.WriteLine("Successfully connected");
            return resp;
        }

        public async Task<Dictionary<string, int>> GetPermissionList()
        {
            var endpoint = Environment.GetEnvironmentVariable("PERMISSIONS_ENDPOINT");
            Console.WriteLine($"Endpoint: {endpoint}");
            var response = await GetDataFromServiceAsync(endpoint+"/permissions");
            Console.WriteLine(response.ToString());
            var responseMessage = JsonConvert.DeserializeObject<PermissionResponse>(response);
            var permDict = new Dictionary<string, int>();
            if (responseMessage != null && responseMessage.Data != null)
            { 
                foreach (var perm in responseMessage.Data)
                {
                    permDict[perm.PermissionName] = perm.PermissionId;
                }
            }
            return permDict;
        }

        public async Task<string> GetPublicKey()
        {
            var endpoint = Environment.GetEnvironmentVariable("SESSIONS_ENDPOINT");
            var response = await GetDataFromServiceAsync(endpoint + "/sessions/public-key");
            Console.WriteLine(response.ToString());
            var responseMessage = JsonConvert.DeserializeObject<PublicKeyResponse>(response);
            if (responseMessage != null && responseMessage.Data != null)
            {
                return responseMessage.Data.PublicKey;
            }
            return "";
        }
    }
}
