public interface IPermissionValidator
{
    bool ValidatePermissions(string permissionName, string sessionKey);
}
