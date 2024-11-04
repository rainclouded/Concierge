namespace task_system_server.Models
{


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
