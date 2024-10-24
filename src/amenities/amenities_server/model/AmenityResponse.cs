namespace amenities_server.model
{
    public class AmenityResponse<T>
    {
        public string Message { get; set; }
        public T Data { get; set; }
        public DateTime Timestamp { get; set; }

        public AmenityResponse(string message, T data)
        {
            Message = message;
            Data = data;
            Timestamp = DateTime.UtcNow;
        }
    }
}