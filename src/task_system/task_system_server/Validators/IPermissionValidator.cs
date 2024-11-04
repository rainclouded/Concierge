namespace task_system_server.Validators
{
    public interface IPermissionValidator
    {
        public abstract bool ValidatePermissions(string permissionName, string sessionkey);
    }
}
