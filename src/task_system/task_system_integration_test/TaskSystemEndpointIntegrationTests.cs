using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using System.Net;
using Xunit;
using task_system_server.Models;
using task_system_server.Persistences;
using task_system_server.Dtos;
using task_system_server.Controllers;
using task_system_server.Repositories;
using Microsoft.AspNetCore.Http;

namespace task_system_server.Tests.Integration
{
    public class TaskSystemEndpointTests : IDisposable
    {
        private readonly DbContextOptions<TaskSystemDbContext> _dbContextOptions;
        private readonly TaskSystemDbContext _context;
        private readonly TaskSystemController _controller;
        private readonly string _databaseName;

        public TaskSystemEndpointTests()
        {
            // Generate a unique database name
            _databaseName = $"TestTaskSystem_{Guid.NewGuid()}";

            // Setup database connection to a test database
            _dbContextOptions = new DbContextOptionsBuilder<TaskSystemDbContext>()
                .UseNpgsql($"Host=127.0.0.1;Port=50017;Username=postgres;Password=sa;Database={_databaseName}")
                .EnableSensitiveDataLogging()
                .Options;

            // Create context and controller
            _context = new TaskSystemDbContext(_dbContextOptions);
            _context.Database.EnsureCreated();

            // Initialize repository with the test context
            var repository = new PostgresTaskSystemRepository(_context);
            _controller = new TaskSystemController(repository);

            // Setup controller context
            _controller.ControllerContext = new ControllerContext
            {
                HttpContext = new DefaultHttpContext()
            };
        }

        [Fact]
        public async Task GetTasks_ReturnsNotFound_WhenNoTasks()
        {
            // Arrange
            var queryObject = new QueryObject();

            // Act
            var actionResult = await _controller.GetTasks(queryObject);
            var notFoundResult = actionResult as NotFoundObjectResult;

            // Assert
            Assert.NotNull(notFoundResult);
            Assert.Equal((int)HttpStatusCode.NotFound, notFoundResult.StatusCode);

            var response = notFoundResult.Value as TaskSystemResponse<string>;
            Assert.NotNull(response);
            Assert.Equal(ResponseMessages.GET_TASKS_FAILED, response.Message);
            Assert.Null(response.Data);
        }

        [Fact]
        public async Task GetTasks_ReturnsOkWithTasks_WhenTasksExist()
        {
            // Arrange
            var testTask = new TaskItem
            {
                TaskType = TaskItemType.RoomCleaning,
                Description = "Clean room 301",
                RoomId = 301,
                RequesterId = 1,
                Status = TaskItemStatus.Pending,
                CreatedAt = DateTime.UtcNow
            };
            await _context.Tasks.AddAsync(testTask);
            await _context.SaveChangesAsync();

            // Act
            var actionResult = await _controller.GetTasks(new QueryObject());
            var okResult = actionResult as OkObjectResult;

            // Assert
            Assert.NotNull(okResult);
            Assert.Equal((int)HttpStatusCode.OK, okResult.StatusCode);

            var response = okResult.Value as TaskSystemResponse<IEnumerable<TaskItem>>;
            Assert.NotNull(response);
            Assert.Equal(ResponseMessages.GET_TASKS_SUCCESS, response.Message);
            var tasks = Assert.Single(response.Data ?? new List<TaskItem>());
            Assert.Equal(TaskItemType.RoomCleaning, tasks.TaskType);
            Assert.Equal(301, tasks.RoomId);
        }

        [Fact]
        public async Task GetTaskById_ReturnsNotFound_WhenTaskDoesNotExist()
        {
            // Act
            var actionResult = await _controller.GetTaskById(999);
            var notFoundResult = actionResult as NotFoundObjectResult;

            // Assert
            Assert.NotNull(notFoundResult);
            Assert.Equal((int)HttpStatusCode.NotFound, notFoundResult.StatusCode);

            var response = notFoundResult.Value as TaskSystemResponse<int>;
            Assert.NotNull(response);
            Assert.Equal(ResponseMessages.GET_TASK_FAILED, response.Message);
        }

        [Fact]
        public async Task AddTask_ReturnsCreated_WithNewTask()
        {
            // Arrange
            var taskDto = new AddTaskDto
            {
                TaskType = TaskItemType.WakeUpCall,
                Description = "Wake up call for room 502",
                RoomId = 502,
                RequesterId = 3
            };

            // Act
            var actionResult = await _controller.AddTask(taskDto);
            var createdResult = actionResult as CreatedAtActionResult;

            // Assert
            Assert.NotNull(createdResult);
            Assert.Equal((int)HttpStatusCode.Created, createdResult.StatusCode);

            var response = createdResult.Value as TaskSystemResponse<TaskItem>;
            Assert.NotNull(response);
            Assert.Equal(ResponseMessages.CREATE_TASK_SUCCESS, response.Message);
            Assert.Equal(taskDto.TaskType, response.Data?.TaskType);
            Assert.Equal(taskDto.RoomId, response.Data?.RoomId);
            Assert.Equal(TaskItemStatus.Pending, response.Data?.Status);
        }

        [Fact]
        public async Task UpdateTask_ReturnsOk_WhenTaskExists()
        {
            // Arrange
            var task = new TaskItem
            {
                TaskType = TaskItemType.FoodDelivery,
                Description = "Deliver breakfast to room 203",
                RoomId = 203,
                RequesterId = 4,
                Status = TaskItemStatus.Pending,
                CreatedAt = DateTime.UtcNow
            };
            await _context.Tasks.AddAsync(task);
            await _context.SaveChangesAsync();

            var updateDto = new UpdateTaskDto
            {
                Description = "Updated: Deliver lunch to room 203",
                Status = TaskItemStatus.InProgress,
                AssigneeId = 5
            };

            // Act
            var actionResult = await _controller.UpdateTask(task.Id, updateDto);
            var okResult = actionResult as OkObjectResult;

            // Assert
            Assert.NotNull(okResult);
            Assert.Equal((int)HttpStatusCode.OK, okResult.StatusCode);

            var response = okResult.Value as TaskSystemResponse<TaskItem>;
            Assert.NotNull(response);
            Assert.Equal(ResponseMessages.UPDATE_TASK_SUCCESS, response.Message);
            Assert.Equal(updateDto.Description, response.Data?.Description);
            Assert.Equal(updateDto.Status, response.Data?.Status);
            Assert.Equal(updateDto.AssigneeId, response.Data?.AssigneeId);
        }

        [Fact]
        public async Task DeleteTask_ReturnsOk_WhenTaskExists()
        {
            // Arrange
            var task = new TaskItem
            {
                TaskType = TaskItemType.LaundryService,
                Description = "Pick up laundry from room 602",
                RoomId = 602,
                RequesterId = 6,
                Status = TaskItemStatus.Pending,
                CreatedAt = DateTime.UtcNow
            };
            await _context.Tasks.AddAsync(task);
            await _context.SaveChangesAsync();

            // Act
            var actionResult = await _controller.DeleteTask(task.Id);
            var okResult = actionResult as OkObjectResult;

            // Assert
            Assert.NotNull(okResult);
            Assert.Equal((int)HttpStatusCode.OK, okResult.StatusCode);

            var response = okResult.Value as TaskSystemResponse<string>;
            Assert.NotNull(response);
            Assert.Equal(ResponseMessages.DELETE_TASK_SUCCESS, response.Message);
            Assert.Null(await _context.Tasks.FindAsync(task.Id));
        }

        [Fact]
        public async Task UpdateTask_ThrowsKeyNotFoundException_WhenTaskDoesNotExist()
        {
            var updateDto = new UpdateTaskDto
            {
                Description = "Updated description",
                Status = TaskItemStatus.InProgress,
                AssigneeId = 5
            };

            await Assert.ThrowsAsync<KeyNotFoundException>(async () =>
            {
                await _controller.UpdateTask(999, updateDto);
            });
        }


        public void Dispose()
        {
            _context.Database.EnsureDeleted();
            _context.Dispose();
        }
    }
}