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
                TaskType = TaskItemType.Maintenance,
                Description = "There is a leak in the bathroom sink that needs urgent attention.",
                RoomId = 101,
                RequesterId = 1,
                AssigneeId = 2,
                Status = TaskItemStatus.InProgress,
                CreatedAt = new DateTime(2024, 10, 10, 10, 30, 0)
            },
            new TaskItem
            {
                Id = _nextId++,
                TaskType = TaskItemType.Maintenance,
                Description = "Some light bulbs are out in the hallway. Please replace them.",
                RoomId = 102,
                RequesterId = 3,
                AssigneeId = 2,
                Status = TaskItemStatus.Pending,
                CreatedAt = new DateTime(2024, 10, 12, 9, 0, 0)
            },
            new TaskItem
            {
                Id = _nextId++,
                TaskType = TaskItemType.RoomCleaning,
                Description = "The conference room needs to be cleaned before the meeting.",
                RoomId = 201,
                RequesterId = 4,
                AssigneeId = 5,
                Status = TaskItemStatus.Completed,
                CreatedAt = new DateTime(2024, 10, 13, 15, 0, 0)
            },
            new TaskItem
            {
                Id = _nextId++,
                TaskType = TaskItemType.FoodDelivery,
                Description = "Deliver breakfast to room 203.",
                RoomId = 203,
                RequesterId = 6,
                AssigneeId = 7,
                Status = TaskItemStatus.Pending,
                CreatedAt = new DateTime(2024, 10, 16, 8, 0, 0)
            },
            new TaskItem
            {
                Id = _nextId++,
                TaskType = TaskItemType.WakeUpCall,
                Description = "Provide a wake-up call at 6:00 AM for room 204.",
                RoomId = 204,
                RequesterId = 8,
                AssigneeId = 9,
                Status = TaskItemStatus.Completed,
                CreatedAt = new DateTime(2024, 10, 16, 6, 0, 0)
            },
            new TaskItem
            {
                Id = _nextId++,
                TaskType = TaskItemType.LaundryService,
                Description = "Pick up laundry from room 205 and deliver it back clean.",
                RoomId = 205,
                RequesterId = 10,
                AssigneeId = 11,
                Status = TaskItemStatus.InProgress,
                CreatedAt = new DateTime(2024, 10, 17, 10, 0, 0)
            },
            new TaskItem
            {
                Id = _nextId++,
                TaskType = TaskItemType.SpaAndMassage,
                Description = "Schedule a massage for the guest in room 206 at 3:00 PM.",
                RoomId = 206,
                RequesterId = 12,
                AssigneeId = 13,
                Status = TaskItemStatus.Pending,
                CreatedAt = new DateTime(2024, 10, 18, 15, 0, 0)
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
            (!query.Status.HasValue || s.Status == query.Status) &&
            (!query.Year.HasValue || s.CreatedAt.Year == query.Year) &&
            (!query.Month.HasValue || s.CreatedAt.Month == query.Month) &&
            (!query.Day.HasValue || s.CreatedAt.Day == query.Day) &&
            (!query.AccountId.HasValue || s.RequesterId == query.AccountId)
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

        existingTask.TaskType = taskDto.TaskType;
        existingTask.Description = taskDto.Description;
        if (taskDto.AssigneeId.HasValue)
        {
            existingTask.AssigneeId = taskDto.AssigneeId.Value;
        }
        existingTask.Status = taskDto.Status;

        return await Task.FromResult(existingTask);
    }

    public async Task<TaskItem?> UpdateAssigneeAsync(int id, int assigneeId)
    {
        var existingTask = _tasks.FirstOrDefault(t => t.Id == id) ?? throw new KeyNotFoundException("Task not found.");

        existingTask.AssigneeId = assigneeId;

        return await Task.FromResult(existingTask);
    }

    public async Task<bool> DeleteTaskAsync(int id)
    {
        var taskToDelete = _tasks.FirstOrDefault(t => t.Id == id) ?? throw new KeyNotFoundException("Task not found.");
        
        _tasks.Remove(taskToDelete);
        
        return await Task.FromResult(true);
    }
}
