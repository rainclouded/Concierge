using task_system_server.Models;

namespace task_system_server.Validators
{
    /*
    Validator class for checking user permissions via PermissionClient
    Args:
      Clients.PermissionClient permissionClient: The PermissionClient instance used to fetch session data
    Returns:
      None
    */
    public class PermissionValidator : IPermissionValidator
    {
        // Private field for PermissionClient instance
        private readonly Clients.PermissionClient _permClient = permissionClient;

        /*
        Validates if the user has the specified permission based on the session key
        Args:
          string permissionName: The name of the permission to validate
          string sessionKey: The session key associated with the current user's session
          out SessionObj sessionData: The session data associated with the current session
        Returns:
          bool: Returns true if the user has the requested permission, otherwise false
        Sets:
          sessionData: The session data extracted from the permission client response
        */
        public bool ValidatePermissions(string permissionName, string sessionKey, out SessionObj sessionData)
        {
            // Get session data from the permission client using the session key
            var permData = _permClient.GetSessionData(sessionKey).GetAwaiter().GetResult();
            sessionData = permData?.Data?.SessionData;

            // Default to empty list if session permission list is null
            var permissionList = sessionData?.SessionPermissionList ?? new List<string>();

            // Check if the permission name exists in the session permissions list
            return permissionList.Contains(permissionName);
        }
    }
}
