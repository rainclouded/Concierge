namespace task_system_server.Validators
{
    public class MockPermissionValidator : IPermissionValidator
    {
        public bool ValidatePermissions(string permissionName, string sessionkey)
        {
            return true;
        }

        public bool ValidateAccountId(int accountId, string sessionKey)
        {
            return true;
        }
    }
}
