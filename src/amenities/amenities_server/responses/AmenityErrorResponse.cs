public class AmenityResponse<T>
{
        public string ErrorMessage { get; set; }
        public T Data { get; set; }
        public DateTime Timestamp { get; set; }

        public AmenityResponse(string message, T data)
        {
            ErrorMessage = message;
            Data = data;
            Timestamp = DateTime.UtcNow;
        }
}