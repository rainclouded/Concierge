namespace task_system_server.Validators
{
    public class PermissionValidator(Clients.PermissionClient permissionClient) : IPermissionValidator
    {
        private readonly Clients.PermissionClient _permClient = permissionClient;

        public bool ValidatePermissions(string permissionName, string sessionKey)
        {
            var permData = _permClient.GetSessionData(sessionKey).GetAwaiter().GetResult();
            var permissionList = permData?.Data?.SessionData?.SessionPermissionList ?? [];

            return permissionList.Contains(permissionName);
        }

        public bool ValidateAccountId(int accountId, string sessionKey)
        {
            var permData = _permClient.GetSessionData(sessionKey).GetAwaiter().GetResult();

            Console.WriteLine(permData?.Data?.SessionData?.AccountId + " " + accountId);
            return permData?.Data?.SessionData?.AccountId == accountId;
        }
    }
}
