using Microsoft.EntityFrameworkCore;
using task_system_server.Dtos;
using task_system_server.Interfaces;
using task_system_server.Models;
using task_system_server.Persistences;

namespace task_system_server.Repositories;

public class PostgresTaskSystemRepository : ITaskSystemRepository
{
    private readonly TaskSystemDbContext _context;
    public PostgresTaskSystemRepository(TaskSystemDbContext context)
    {
        _context = context;
    }

    public async Task<IEnumerable<TaskItem>> GetTasksAsync(QueryObject query)
    {
        var tasks = _context.Tasks.AsQueryable();
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

        return await tasks.ToListAsync();
    }

    public async Task<TaskItem?> GetTaskByIdAsync(int taskId)
    {
        return await _context.Tasks.FirstOrDefaultAsync(t => t.Id == taskId);
    }

    public async Task<TaskItem> AddTaskAsync(TaskItem task)
    {
        if (_context.Tasks.FirstOrDefault(t => t.Id == task.Id) != null)
        {
            throw new InvalidOperationException("TaskItem with the same ID already exists.");
        }

        //Necessary format to store in DB
        task.CreatedAt = DateTime.UtcNow;

        await _context.Tasks.AddAsync(task);
        await _context.SaveChangesAsync();

        return task;
    }

    public async Task<TaskItem?> UpdateTaskAsync(int id, UpdateTaskDto taskDto)
    {
        var existingTask = await _context.Tasks.FirstOrDefaultAsync(t => t.Id == id) ?? throw new KeyNotFoundException("Task not found.");

        existingTask.TaskType = taskDto.TaskType;
        existingTask.Description = taskDto.Description;
        if (taskDto.AssigneeId.HasValue)
        {
            existingTask.AssigneeId = taskDto.AssigneeId.Value;
        }
        existingTask.Status = taskDto.Status;

        await _context.SaveChangesAsync();

        return existingTask;
    }

    public async Task<TaskItem?> UpdateAssigneeAsync(int id, int assigneeId)
    {
        var existingTask = await _context.Tasks.FirstOrDefaultAsync(t => t.Id == id) ?? throw new KeyNotFoundException("Task not found.");

        existingTask.AssigneeId = assigneeId;

        await _context.SaveChangesAsync();

        return existingTask;
    }

    public async Task<bool> DeleteTaskAsync(int id)
    {
        var taskToDelete = await _context.Tasks.FirstOrDefaultAsync(t => t.Id == id) ?? throw new KeyNotFoundException("Task not found.");

        _context.Tasks.Remove(taskToDelete);
        await _context.SaveChangesAsync();

        return true;
    }
}
