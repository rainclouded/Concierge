using task_system_server.Models;
using task_system_server.Interfaces;
using task_system_server.Dtos;

namespace task_system_server.Repositories;

public class StubTaskSystemRepository : ITaskSystemRepository
{
    private readonly List<TaskItem> _tasks;
    private int _nextId = 1;

    public StubTaskSystemRepository()
    {
        _tasks =
        [
            new TaskItem
            {
                Id = _nextId++,
                Title = "Fix the bathroom leak",
                Description = "There is a leak in the bathroom sink that needs urgent attention.",
                RoomId = 101,
                RequesterId = 1,
                AssigneeId = 2,
                Status = "In Progress",
                CreatedAt = new DateTime(2024, 10, 10, 10, 30, 0)
            },
            new TaskItem
            {
                Id = _nextId++,
                Title = "Replace light bulbs in the hallway",
                Description = "Some light bulbs are out in the hallway. Please replace them.",
                RoomId = 102,
                RequesterId = 3,
                AssigneeId = 2,
                Status = "Pending",
                CreatedAt = new DateTime(2024, 10, 12, 9, 0, 0)
            },
            new TaskItem
            {
                Id = _nextId++,
                Title = "Clean the conference room",
                Description = "The conference room needs to be cleaned before the meeting.",
                RoomId = 201,
                RequesterId = 4,
                AssigneeId = 5,
                Status = "Completed",
                CreatedAt = new DateTime(2024, 10, 13, 15, 0, 0)
            },
            new TaskItem
            {
                Id = _nextId++,
                Title = "Check fire alarm batteries",
                Description = "Ensure that all fire alarms have functioning batteries.",
                RoomId = 103,
                RequesterId = 2,
                AssigneeId = 3,
                Status = "Pending",
                CreatedAt = new DateTime(2024, 10, 15, 11, 0, 0)
            },
            new TaskItem
            {
                Id = _nextId++,
                Title = "Organize supplies in storage",
                Description = "Organize the storage area to make supplies easily accessible.",
                RoomId = 104,
                RequesterId = 1,
                AssigneeId = 4,
                Status = "In Progress",
                CreatedAt = new DateTime(2024, 10, 18, 14, 0, 0)
            }
        ];
    }

    public async Task<IEnumerable<TaskItem>> GetTasksAsync(QueryObject query)
    {
        var tasks = _tasks.AsQueryable();
        tasks = tasks.Where(s =>
            (!query.RoomId.HasValue || s.RoomId == query.RoomId) &&
            (!query.RequesterId.HasValue || s.RequesterId == query.RequesterId) &&
            (!query.AssigneeId.HasValue || s.AssigneeId == query.AssigneeId) &&
            (string.IsNullOrWhiteSpace(query.Status) || s.Status.Equals(query.Status)) &&
            (!query.Year.HasValue || s.CreatedAt.Year == query.Year) &&
            (!query.Month.HasValue || s.CreatedAt.Month == query.Month) &&
            (!query.Day.HasValue || s.CreatedAt.Day == query.Day)
        );

        tasks = query.SortAscending ? tasks.OrderBy(s => s.CreatedAt) : tasks.OrderByDescending(s => s.CreatedAt);

        return await Task.FromResult(tasks);
    }

    public async Task<TaskItem?> GetTaskByIdAsync(int id)
    {
        return await Task.FromResult(_tasks.FirstOrDefault(t => t.Id == id));
    }

    public async Task<TaskItem> AddTaskAsync(TaskItem task)
    {
        if (_tasks.FirstOrDefault(t => t.Id == task.Id) != null)
        {
            throw new InvalidOperationException("TaskItem with the same ID already exists.");
        }

        task.Id = _nextId++;
        _tasks.Add(task);

        return await Task.FromResult(task);
    }

    public async Task<TaskItem?> UpdateTaskAsync(int id, UpdateTaskDto taskDto)
    {
        var existingTask = _tasks.FirstOrDefault(t => t.Id == id) ?? throw new KeyNotFoundException("Task not found.");

        existingTask.Title = taskDto.Title;
        existingTask.Description = taskDto.Description;
        existingTask.AssigneeId = taskDto.AssigneeId;
        existingTask.Status = taskDto.Status;

        return await Task.FromResult(existingTask);
    }

    public async Task<bool> DeleteTaskAsync(int id)
    {
        var taskToDelete = _tasks.FirstOrDefault(t => t.Id == id) ?? throw new KeyNotFoundException("Task not found.");
        
        _tasks.Remove(taskToDelete);
        
        return await Task.FromResult(true);
    }
}
