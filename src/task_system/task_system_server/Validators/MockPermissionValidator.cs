namespace task_system_server.Validators
{
    public class MockPermissionValidator : IPermissionValidator
    {
        public bool ValidatePermissions(string permissionName, string sessionkey)
        {
            return true;
        }
    }
}
