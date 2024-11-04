namespace task_system_server.Exceptions
{
    public class ClientConnectionException : Exception
    {
        public ClientConnectionException(string message, Exception innerException) : base(message)
        {
        }

        public ClientConnectionException(string message) : base(message)
        {
        }

        public ClientConnectionException()
        {
        }

        public int ErrorCode { get; set; }
    }
}
