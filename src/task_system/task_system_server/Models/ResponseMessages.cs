namespace task_system_server.Models;

public class ResponseMessages
{
    public const string GET_TASKS_SUCCESS = "Tasks retrieved successfully.";
    public const string GET_TASKS_FAILED = "No task was found.";
    public const string GET_TASK_SUCCESS = "Task retrieved successfully";
    public const string GET_TASK_FAILED = "Task with specified id not found.";
    public const string CREATE_TASK_SUCCESS = "Task created successfully.";
    public const string UPDATE_TASK_SUCCESS = "Task updated successfully.";
    public const string DELETE_TASK_SUCCESS = "Task deleted successfully.";
    public const string INVALID_TASK_PASSED = "Bad request. Task with invalid parameters was passed.";
}
