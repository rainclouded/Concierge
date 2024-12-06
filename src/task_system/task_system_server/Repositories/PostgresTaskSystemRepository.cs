using Microsoft.EntityFrameworkCore;
using task_system_server.Dtos;
using task_system_server.Interfaces;
using task_system_server.Models;
using task_system_server.Persistences;

namespace task_system_server.Repositories;

public class PostgresTaskSystemRepository : ITaskSystemRepository
{
    private readonly TaskSystemDbContext _context;

    /*
    Constructor for PostgresTaskSystemRepository
    Args:
      TaskSystemDbContext context: The context used to interact with the database
    Returns:
      None
    */
    public PostgresTaskSystemRepository(TaskSystemDbContext context)
    {
        _context = context;
    }

    /*
    Retrieves tasks based on the provided query filters
    Args:
      QueryObject query: Object containing various filter and sorting parameters
    Returns:
      Task<IEnumerable<TaskItem>>: A task representing the asynchronous operation that returns the filtered list of tasks
    */
    public async Task<IEnumerable<TaskItem>> GetTasksAsync(QueryObject query)
    {
        var tasks = _context.Tasks.AsQueryable();
        
        // Apply filters based on the query parameters
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

        // Apply sorting based on the query parameter
        tasks = query.SortAscending ? tasks.OrderBy(s => s.CreatedAt) : tasks.OrderByDescending(s => s.CreatedAt);

        return await tasks.ToListAsync();
    }

    /*
    Retrieves a task by its ID
    Args:
      int taskId: The ID of the task to retrieve
    Returns:
      Task<TaskItem?>: A task representing the asynchronous operation that returns the task or null if not found
    */
    public async Task<TaskItem?> GetTaskByIdAsync(int taskId)
    {
        return await _context.Tasks.FirstOrDefaultAsync(t => t.Id == taskId);
    }

    /*
    Adds a new task to the database
    Args:
      TaskItem task: The task to add
    Returns:
      Task<TaskItem>: A task representing the asynchronous operation that returns the added task
    Throws:
      InvalidOperationException: If a task with the same ID already exists
    */
    public async Task<TaskItem> AddTaskAsync(TaskItem task)
    {
        // Check if task with the same ID already exists
        if (_context.Tasks.FirstOrDefault(t => t.Id == task.Id) != null)
        {
            throw new InvalidOperationException("TaskItem with the same ID already exists.");
        }

        // Necessary format to store in DB
        task.CreatedAt = DateTime.UtcNow;

        await _context.Tasks.AddAsync(task);
        await _context.SaveChangesAsync();

        return task;
    }

    /*
    Updates an existing task with new values
    Args:
      int id: The ID of the task to update
      UpdateTaskDto taskDto: Object containing the updated task values
    Returns:
      Task<TaskItem?>: A task representing the asynchronous operation that returns the updated task
    Throws:
      KeyNotFoundException: If the task with the given ID does not exist
    */
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

    /*
    Updates the assignee of an existing task
    Args:
      int id: The ID of the task to update
      int assigneeId: The ID of the new assignee
    Returns:
      Task<TaskItem?>: A task representing the asynchronous operation that returns the updated task
    Throws:
      KeyNotFoundException: If the task with the given ID does not exist
    */
    public async Task<TaskItem?> UpdateAssigneeAsync(int id, int assigneeId)
    {
        var existingTask = await _context.Tasks.FirstOrDefaultAsync(t => t.Id == id) ?? throw new KeyNotFoundException("Task not found.");

        existingTask.AssigneeId = assigneeId;

        await _context.SaveChangesAsync();

        return existingTask;
    }

    /*
    Deletes a task from the database
    Args:
      int id: The ID of the task to delete
    Returns:
      Task<bool>: A task representing the asynchronous operation that returns true if the task was deleted
    Throws:
      KeyNotFoundException: If the task with the given ID does not exist
    */
    public async Task<bool> DeleteTaskAsync(int id)
    {
        var taskToDelete = await _context.Tasks.FirstOrDefaultAsync(t => t.Id == id) ?? throw new KeyNotFoundException("Task not found.");

        _context.Tasks.Remove(taskToDelete);
        await _context.SaveChangesAsync();

        return true;
    }
}
