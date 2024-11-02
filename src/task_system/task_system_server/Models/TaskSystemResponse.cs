namespace task_system_server.Models;

public class TaskSystemResponse<T>(string message, T? data)
{
    public string Message { get; set; } = message;
    public T? Data { get; set; } = data;
    public DateTime Timestamp { get; set; } = DateTime.UtcNow;
}
