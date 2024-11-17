using task_system_server.Dtos;
using task_system_server.Models;

namespace task_system_server.Interfaces;

public interface ITaskSystemRepository
{
    Task<IEnumerable<TaskItem>> GetTasksAsync(QueryObject query);
    Task<TaskItem?> GetTaskByIdAsync(int id);
    Task<TaskItem> AddTaskAsync(TaskItem task);
    Task<TaskItem?> UpdateTaskAsync(int id, UpdateTaskDto task);
    Task<TaskItem?> UpdateAssigneeAsync(int id, int assigneeId);
    Task<bool> DeleteTaskAsync(int id);
}
