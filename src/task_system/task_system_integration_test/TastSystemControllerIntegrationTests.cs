using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Microsoft.AspNetCore.Http;
using Xunit;
using task_system_server.Models;
using task_system_server.Persistences;
using task_system_server.Dtos;
using task_system_server.Controllers;
using task_system_server.Repositories;
using task_system_server.Validators;

namespace task_system_server.Tests.Integration
{
    public class TaskSystemControllerTests : IDisposable
    {
        private readonly DbContextOptions<TaskSystemDbContext> _dbContextOptions;
        private readonly TaskSystemDbContext _context;
        private readonly TaskSystemController _controller;
        private readonly string _databaseName;

        public TaskSystemControllerTests()
        {
            // Generate a unique database name
            _databaseName = $"TestTaskSystem_{Guid.NewGuid()}";

            // Setup database connection to a test database
            _dbContextOptions = new DbContextOptionsBuilder<TaskSystemDbContext>()
                .UseNpgsql($"Host=127.0.0.1;Port=50021;Username=postgres;Password=sa;Database={_databaseName}")
                .EnableSensitiveDataLogging()
                .Options;

            // Create context and controller
            _context = new TaskSystemDbContext(_dbContextOptions);
            _context.Database.EnsureCreated(); // Create the test database

            var httpContext = new DefaultHttpContext();
            httpContext.Request.Headers["X-API-Key"] = "TestsKey";

            // Initialize repository with the test context
            var repository = new PostgresTaskSystemRepository(_context);
            _controller = new TaskSystemController(repository, new MockPermissionValidator())
            {
                ControllerContext = new ControllerContext
                {
                    HttpContext = httpContext
                }
            };
        }

        [Fact]
        public async Task GetTasks_ReturnsNotFound_WhenNoTasks()
        {
            // Act
            var result = await _controller.GetTasks(new QueryObject());

            // Assert
            var notFoundResult = Assert.IsType<NotFoundObjectResult>(result);
            var response = Assert.IsType<TaskSystemResponse<string>>(notFoundResult.Value);
            Assert.Equal(ResponseMessages.GET_TASKS_FAILED, response.Message);
            Assert.Null(response.Data);
        }

        [Fact]
        public async Task GetTasks_ReturnsOkResult_WithTasks()
        {
            // Arrange
            var testTask = new TaskItem
            {
                TaskType = TaskItemType.RoomCleaning,
                Description = "Clean room 301",
                RoomId = 301,
                RequesterId = 1,
                Status = TaskItemStatus.Pending,
                CreatedAt = DateTime.Now
            };
            await _context.Tasks.AddAsync(testTask);
            await _context.SaveChangesAsync();

            // Act
            var result = await _controller.GetTasks(new QueryObject());

            // Assert
            var okResult = Assert.IsType<OkObjectResult>(result);
            var response = Assert.IsType<TaskSystemResponse<IEnumerable<TaskItem>>>(okResult.Value);
            Assert.Equal(ResponseMessages.GET_TASKS_SUCCESS, response.Message);
            Assert.NotNull(response.Data); // Ensure Data is not null
            var tasks = Assert.Single(response.Data ?? new List<TaskItem>());
            Assert.Equal(TaskItemType.RoomCleaning, tasks.TaskType);
            Assert.Equal(301, tasks.RoomId);
        }

        [Fact]
        public async Task GetTaskById_ReturnsNotFound_ForNonexistentTask()
        {
            // Act
            var result = await _controller.GetTaskById(999);

            // Assert
            var notFoundResult = Assert.IsType<NotFoundObjectResult>(result);
            var response = Assert.IsType<TaskSystemResponse<int>>(notFoundResult.Value);
            Assert.Equal(ResponseMessages.GET_TASK_FAILED, response.Message);
            Assert.Equal(999, response.Data);
        }

        [Fact]
        public async Task GetTaskById_ReturnsOk_ForExistingTask()
        {
            // Arrange
            var testTask = new TaskItem
            {
                TaskType = TaskItemType.Maintenance,
                Description = "Fix AC in room 405",
                RoomId = 405,
                RequesterId = 2,
                Status = TaskItemStatus.Pending,
                CreatedAt = DateTime.Now
            };
            await _context.Tasks.AddAsync(testTask);
            await _context.SaveChangesAsync();

            // Act
            var result = await _controller.GetTaskById(testTask.Id);

            // Assert
            var okResult = Assert.IsType<OkObjectResult>(result);
            var response = Assert.IsType<TaskSystemResponse<TaskItem>>(okResult.Value);
            Assert.Equal(ResponseMessages.GET_TASK_SUCCESS, response.Message);
            Assert.NotNull(response.Data);
            Assert.Equal(testTask.TaskType, response.Data.TaskType);
            Assert.Equal(testTask.RoomId, response.Data.RoomId);
            Assert.Equal(testTask.Description, response.Data.Description);
        }

        [Fact]
        public async Task AddTask_ReturnsCreatedAtAction()
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
            var result = await _controller.AddTask(taskDto);

            // Assert
            var createdAtActionResult = Assert.IsType<CreatedAtActionResult>(result);
            var response = Assert.IsType<TaskSystemResponse<TaskItem>>(createdAtActionResult.Value);
            Assert.Equal(ResponseMessages.CREATE_TASK_SUCCESS, response.Message);
            Assert.NotNull(response.Data);
            Assert.Equal(taskDto.TaskType, response.Data.TaskType);
            Assert.Equal(taskDto.RoomId, response.Data.RoomId);
            Assert.Equal(taskDto.Description, response.Data.Description);
            Assert.Equal(TaskItemStatus.Pending, response.Data.Status);
        }

        [Fact]
        public async Task UpdateTask_ReturnsOkResult_WithUpdatedTask()
        {
            // Arrange
            var task = new TaskItem
            {
                TaskType = TaskItemType.FoodDelivery,
                Description = "Deliver breakfast to room 203",
                RoomId = 203,
                RequesterId = 4,
                Status = TaskItemStatus.Pending,
                CreatedAt = DateTime.Now
            };
            await _context.Tasks.AddAsync(task);
            await _context.SaveChangesAsync();

            var updateDto = new UpdateTaskDto
            {
                Description = "Deliver lunch to room 203",
                Status = TaskItemStatus.InProgress,
                AssigneeId = 5
            };

            // Act
            var result = await _controller.UpdateTask(task.Id, updateDto);

            // Assert
            var okResult = Assert.IsType<OkObjectResult>(result);
            var response = Assert.IsType<TaskSystemResponse<TaskItem>>(okResult.Value);
            Assert.Equal(ResponseMessages.UPDATE_TASK_SUCCESS, response.Message);
            Assert.NotNull(response.Data);
            Assert.Equal(updateDto.Description, response.Data.Description);
            Assert.Equal(updateDto.Status, response.Data.Status);
            Assert.Equal(updateDto.AssigneeId, response.Data.AssigneeId);
        }

        [Fact]
        public async Task UpdateAssignee_ReturnsOkResult_WithUpdatedAssigneeId()
        {
            // Arrange
            var task = new TaskItem
            {
                TaskType = TaskItemType.SpaAndMassage,
                Description = "Schedule a massage for room 305",
                RoomId = 305,
                RequesterId = 7,
                AssigneeId = 4,
                Status = TaskItemStatus.Pending,
                CreatedAt = DateTime.Now
            };
            await _context.Tasks.AddAsync(task);
            await _context.SaveChangesAsync();

            var updateAssigneeDto = new UpdateAssigneeDto
            {
                AssigneeId = 10
            };

            // Act
            var result = await _controller.UpdateAssignee(task.Id, updateAssigneeDto);

            // Assert
            var okResult = Assert.IsType<OkObjectResult>(result);
            var response = Assert.IsType<TaskSystemResponse<TaskItem>>(okResult.Value);
            Assert.Equal(ResponseMessages.UPDATE_TASK_SUCCESS, response.Message);
            Assert.NotNull(response.Data);
            Assert.Equal(updateAssigneeDto.AssigneeId, response.Data.AssigneeId);
            Assert.Equal(task.TaskType, response.Data.TaskType);
            Assert.Equal(task.Description, response.Data.Description);
            Assert.Equal(task.Status, response.Data.Status);
            Assert.Equal(task.RoomId, response.Data.RoomId);
        }

        [Fact]
        public async Task UpdateAssignee_ThrowsKeyNotFoundException_ForNonexistentTask()
        {
            // Arrange
            var updateAssigneeDto = new UpdateAssigneeDto
            {
                AssigneeId = 10
            };

            // Act & Assert
            await Assert.ThrowsAsync<KeyNotFoundException>(async () =>
            {
                await _controller.UpdateAssignee(999, updateAssigneeDto);
            });
        }



        [Fact]
        public async Task DeleteTask_ReturnsNotFound_WhenTaskDoesNotExist()
        {
            // Act
            var result = await _controller.DeleteTask(999);

            // Assert
            var notFoundResult = Assert.IsType<NotFoundObjectResult>(result);
            var response = Assert.IsType<TaskSystemResponse<int>>(notFoundResult.Value);
            Assert.Equal(ResponseMessages.GET_TASK_FAILED, response.Message);
        }

        [Fact]
        public async Task DeleteTask_ReturnsOkResult_WhenTaskExists()
        {
            // Arrange
            var task = new TaskItem
            {
                TaskType = TaskItemType.LaundryService,
                Description = "Pick up laundry from room 602",
                RoomId = 602,
                RequesterId = 6,
                Status = TaskItemStatus.Pending,
                CreatedAt = DateTime.Now
            };
            await _context.Tasks.AddAsync(task);
            await _context.SaveChangesAsync();

            // Act
            var result = await _controller.DeleteTask(task.Id);

            // Assert
            var okResult = Assert.IsType<OkObjectResult>(result);
            var response = Assert.IsType<TaskSystemResponse<string>>(okResult.Value);
            Assert.Equal(ResponseMessages.DELETE_TASK_SUCCESS, response.Message);
            Assert.Null(response.Data);
            Assert.Null(await _context.Tasks.FindAsync(task.Id));
        }

        public void Dispose()
        {
            // Clean up the test database after each test
            _context.Database.EnsureDeleted();
            _context.Dispose();
        }
    }
}