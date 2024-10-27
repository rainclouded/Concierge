using System.Linq.Expressions;
using Microsoft.EntityFrameworkCore;
using Moq;
using task_system_server.Dtos;
using task_system_server.Models;
using task_system_server.Persistences;
using task_system_server.Repositories;

namespace task_system_test;

public class PostgresTaskSystemRepositoryTest : IDisposable
{
     private readonly TaskSystemDbContext _context;
    private readonly PostgresTaskSystemRepository _repository;

    public PostgresTaskSystemRepositoryTest()
    {
        var options = new DbContextOptionsBuilder<TaskSystemDbContext>()
            .UseInMemoryDatabase(databaseName: "TaskSystemTestDb")
            .Options;

        _context = new TaskSystemDbContext(options);
        _repository = new PostgresTaskSystemRepository(_context);

        SeedDatabase();
    }

    private void SeedDatabase()
    {
        // Seed the in-memory database with test data
        _context.Tasks.AddRange(new List<TaskItem>
        {
            new TaskItem
            {
                Id = 1,
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
                Id = 2,
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
                Id = 3,
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
                Id = 4,
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
                Id = 5,
                Title = "Organize supplies in storage",
                Description = "Organize the storage area to make supplies easily accessible.",
                RoomId = 104,
                RequesterId = 1,
                AssigneeId = 4,
                Status = "In Progress",
                CreatedAt = new DateTime(2024, 10, 18, 14, 0, 0)
            }
        });
        _context.SaveChanges();
    }

    [Fact]
    public async Task GetTaskAsync_NoFilter_ReturnsAllTasks()
    {
        var query = new QueryObject();

        var result = await _repository.GetTasksAsync(query);

        Assert.Equal(5, result.Count());
    }

    [Theory]
    [InlineData(101)]
    [InlineData(102)]
    public async Task GetTaskAsync_FilterByRoomId_ReturnsMatchingTasks(int roomId)
    {
        var query = new QueryObject { RoomId = roomId };

        var result = await _repository.GetTasksAsync(query);

        Assert.All(result, task => Assert.Equal(roomId, task.RoomId));
    }

    [Theory]
    [InlineData(2)]
    [InlineData(3)]
    public async Task GetTaskAsync_FilterByAssignedId_ReturnsMatchingTasks(int assigned)
    {
        var query = new QueryObject { AssigneeId = assigned };

        var result = await _repository.GetTasksAsync(query);

        Assert.All(result, task => Assert.Equal(assigned, task.AssigneeId));
    }

    [Theory]
    [InlineData("Pending")]
    [InlineData("In Progress")]
    [InlineData("Completed")]
    public async Task GetTaskAsync_FilterByStatus_ReturnsMatchingTasks(string status)
    {
        var query = new QueryObject { Status = status };

        var result = await _repository.GetTasksAsync(query);

        Assert.All(result, task => Assert.Equal(status, task.Status));
    }

    [Fact]
    public async Task GetTaskAsync_FilterByDate_ReturnsMatchingTasks()
    {
        var query = new QueryObject { Year = 2024, Month = 10, Day = 10 };

        var result = await _repository.GetTasksAsync(query);

        Assert.All(result, task => Assert.Equal(2024, task.CreatedAt.Year));
        Assert.All(result, task => Assert.Equal(10, task.CreatedAt.Month));
        Assert.All(result, task => Assert.Equal(10, task.CreatedAt.Day));
    }

    [Theory]
    [InlineData(true)]
    [InlineData(false)]
    public async Task GetTaskAsync_SortByDate_ReturnsSortedTasks(bool ascending)
    {
        var query = new QueryObject { SortAscending = ascending };

        var result = (await _repository.GetTasksAsync(query)).ToList();

        for (int i = 1; i < result.Count; i++)
        {
            if (ascending)
            {
                Assert.True(result[i - 1].CreatedAt <= result[i].CreatedAt);
            }
            else
            {
                Assert.True(result[i - 1].CreatedAt >= result[i].CreatedAt);
            }
        }
    }

    [Fact]
    public async Task GetTaskByIdAsync_ExistingId_ReturnsTask()
    {
        int taskId = 1;

        var result = await _repository.GetTaskByIdAsync(taskId);

        Assert.NotNull(result);
        Assert.Equal(taskId, result.Id);
    }

    [Fact]
    public async Task GetTaskByIdAsync_NonexistingId_ReturnsNull()
    {
        int taskId = 999;

        var result = await _repository.GetTaskByIdAsync(taskId);

        Assert.Null(result);
    }

    [Fact]
    public async Task AddTask_ValidTask_AddsTaskAndReturnsWithNewId()
    {
        var newTask = new TaskItem
        {
            Title = "New Task",
            Description = "Test Description",
            RoomId = 105,
            RequesterId = 1,
            AssigneeId = null,
            Status = "Pending",
            CreatedAt = DateTime.Now
        };

        var result = await _repository.AddTaskAsync(newTask);
        var addedTask = await _repository.GetTaskByIdAsync(result.Id);

        Assert.NotNull(addedTask);
        Assert.Equal(newTask.Title, addedTask.Title);
        Assert.Equal(newTask.Description, addedTask.Description);
        Assert.Equal(newTask.Description, addedTask.Description);
        Assert.Equal(newTask.RoomId, addedTask.RoomId);
        Assert.Equal(newTask.RequesterId, addedTask.RequesterId);
        Assert.True(addedTask.Id > 0);
    }

    [Fact]
    public async Task AddTasksAsync_Duplicated_ThrowsInvalidOperationException()
    {
        var existingTask = await _repository.GetTaskByIdAsync(1);
        var duplicatedTask = new TaskItem
        {
            Id = 1,
            Title = "New Task",
            Description = "Test Description",
            RoomId = 105,
            RequesterId = 1,
        };

        await Assert.ThrowsAsync<InvalidOperationException>(() => _repository.AddTaskAsync(duplicatedTask));
    }

    [Fact]
    public async Task UpdateTaskAsync_ExistingTask_UpdatesAndReturnsTask()
    {
        int taskId = 1;
        var updateDto = new UpdateTaskDto
        {
            Title = "Updated Title",
            Description = "Updated Description",
            AssigneeId = 3,
            Status = "Completed"
        };

        var result = await _repository.UpdateTaskAsync(taskId, updateDto);

        Assert.NotNull(result);
        Assert.Equal(updateDto.Title, result.Title);
        Assert.Equal(updateDto.Description, result.Description);
        Assert.Equal(updateDto.AssigneeId, result.AssigneeId);
        Assert.Equal(updateDto.Status, result.Status);
    }

    [Fact]
    public async Task UpdateTaskAsync_NonexistentTask_ThrowsKeyNotFoundException()
    {
        int taskId = 999;
        var updateDto = new UpdateTaskDto
        {
            Title = "Updated Title",
            Description = "Updated Description",
            AssigneeId = 3,
            Status = "Completed"
        };

        await Assert.ThrowsAsync<KeyNotFoundException>(() => _repository.UpdateTaskAsync(taskId, updateDto));
    }

    [Fact]
    public async Task DeleteTaskAsync_ExistingTask_DeletesAndReturnsTrue()
    {
        int taskId = 1;

        var result = await _repository.DeleteTaskAsync(taskId);
        var deletedTask = await _repository.GetTaskByIdAsync(taskId);

        Assert.True(result);
        Assert.Null(deletedTask);
    }

    [Fact]
    public async Task DeleteTaskAsync_NonexistingTask_ThrowsKeyNotFoundException()
    {
        int taskId = 999;

        await Assert.ThrowsAsync<KeyNotFoundException>(() => _repository.DeleteTaskAsync(taskId));
    }

    public void Dispose()
    {
        // Ensures the in-memory database is cleared after each test run
        _context.Database.EnsureDeleted();
        _context.Dispose();
        // Prevents finalizer from being called, improving performance and preventing derived classes from needing to reimplement Dispose
        GC.SuppressFinalize(this);
    }
}
