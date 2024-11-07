namespace amenities_server.validators
{
    public class MockPermissionValidator : IPermissionValidator
    {
        public bool ValidatePermissions(string permissionName, string sessionKey)
        {
            return true;
        }
    }
}
