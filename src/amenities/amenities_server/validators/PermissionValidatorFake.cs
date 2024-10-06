public class PermissionValidatorFake : IPermissionValidator
{
    public bool ValidatePermissions(string permissionName, string sessionKey)
    {
        return true;
    }
}