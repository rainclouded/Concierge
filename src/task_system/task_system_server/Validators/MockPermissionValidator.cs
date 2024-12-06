using task_system_server.Models;

namespace task_system_server.Validators
{
    /*
    Mock implementation of IPermissionValidator for simulating permission validation
    Args:
      None
    Returns:
      None
    */
    public class MockPermissionValidator : IPermissionValidator
    {
        /*
        Validates permissions based on a given permission name and session key
        Args:
          string permissionName: The name of the permission to validate
          string sessionkey: The session key associated with the current user
          out SessionObj sessionData: The session data returned after validation
        Returns:
          bool: Returns true if permissions are valid, false otherwise
        Sets:
          sessionData: A SessionObj containing sample account details and associated permissions
        */
        public bool ValidatePermissions(string permissionName, string sessionkey, out SessionObj sessionData)
        {
            // Simulating session data with pre-set values
            sessionData = new SessionObj
            {
                AccountId = 4,
                AccountName = "0",
                SessionPermissionList = new List<string> { "VIEW_TASKS", "EDIT_TASKS" }
            };

            return true; // Simulates successful permission validation
        }
    }
}
