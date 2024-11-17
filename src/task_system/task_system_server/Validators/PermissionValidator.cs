using task_system_server.Models;

namespace task_system_server.Validators
{
    public class PermissionValidator(Clients.PermissionClient permissionClient) : IPermissionValidator
    {
        private readonly Clients.PermissionClient _permClient = permissionClient;

        public bool ValidatePermissions(string permissionName, string sessionKey, out SessionObj sessionData)
        {
            var permData = _permClient.GetSessionData(sessionKey).GetAwaiter().GetResult();
            sessionData = permData?.Data?.SessionData;
            var permissionList = sessionData?.SessionPermissionList ?? [];

            return permissionList.Contains(permissionName);
        }
    }
}
