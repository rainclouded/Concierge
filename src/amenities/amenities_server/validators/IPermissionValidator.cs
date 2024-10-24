namespace amenities_server.validators
{
    public interface IPermissionValidator
    {
        bool ValidatePermissions(string permissionName, string sessionKey);
    }
}