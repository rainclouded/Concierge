namespace amenities_server.exceptions
{
    public class ConnectionError: Exception
    {
        public ConnectionError(string message, Exception innerException) : base(message)
        {
        }

        public ConnectionError(string message) : base(message)
        { 
        }

        public ConnectionError()
        {
        }

        public int ErrorCode { get; set; }
    }
}
