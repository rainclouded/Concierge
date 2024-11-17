using task_system_server.Models;

namespace task_system_server.Validators
{
    public class MockPermissionValidator : IPermissionValidator
    {
        public bool ValidatePermissions(string permissionName, string sessionkey, out SessionObj sessionData)
        {
            sessionData = new SessionObj
            {
                AccountId = 4,
                AccountName = "0",
                SessionPermissionList = new List<string> { "VIEW_TASKS", "EDIT_TASKS" }
            };

            return true;
        }
    }
}
