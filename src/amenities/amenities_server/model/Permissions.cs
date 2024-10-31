namespace amenities_server.model
{
    public class Permission
    {
        public int PermissionId { get; set; }
        public required string PermissionName { get; set; }
        public bool PermissionState { get; set; }
    }

    public class PermissionResponse
    {
        public required string Message { get; set; }
        public List<Permission>? Data { get; set; }
        public DateTime Timestamp { get; set; }
    }

    public class PublicKeyResponse
    {
        public required string Message { get; set; }
        public PublicKeyData? Data { get; set; }
        public DateTime Timestamp { get; set; }
    }

    public class PublicKeyData
    {
        public required string PublicKey { get; set; }
    }

    public class SessionData
    {
        public required string Message { get; set; }
        public required SessionDataDict Data { get; set; }
        public DateTime Timestamp { get; set; }
    }

    public class SessionDataDict
    {
        public required SessionObj SessionData { get; set; }
    }

    public class SessionObj
    {
        public int AccountId { get; set; }
        public required string AccountName { get; set; }
        public required List<string> SessionPermissionList { get; set; }
    }
}