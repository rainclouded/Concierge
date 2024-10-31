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
}