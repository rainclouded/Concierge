using task_system_server.Models;

namespace task_system_server.Validators
{
    public interface IPermissionValidator
    {
        public abstract bool ValidatePermissions(string permissionName, string sessionkey, out SessionObj sessionData);
    }
}
