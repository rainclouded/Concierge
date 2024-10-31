using Microsoft.AspNetCore.Mvc;
using Moq;
using task_system_server.Controllers;
using task_system_server.Dtos;
using task_system_server.Interfaces;
using task_system_server.Models;

namespace task_system_test
{
    public class TaskSystemControllerTest
    {
        private readonly Mock<ITaskSystemRepository> _mockRepo;
        private readonly TaskSystemController _controller;

        public TaskSystemControllerTest()
        {
            _mockRepo = new Mock<ITaskSystemRepository>();
            _controller = new TaskSystemController(_mockRepo.Object);
        }

        [Fact]
        public async Task GetTasks_ReturnsOKResult_WhenTasksExist()
        {
            var query = new QueryObject();
            var tasks = new List<TaskItem>{
                new TaskItem
                {
                    Id = 1,
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
                    Id = 2,
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
                    Id = 3,
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
                    Id = 4,
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
                    Id = 5,
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
                    Id = 6,
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
                    Id = 7,
                    TaskType = TaskItemType.SpaAndMassage,
                    Description = "Schedule a massage for the guest in room 206 at 3:00 PM.",
                    RoomId = 206,
                    RequesterId = 12,
                    AssigneeId = 13,
                    Status = TaskItemStatus.Pending,
                    CreatedAt = new DateTime(2024, 10, 18, 15, 0, 0)
                }
            };

            _mockRepo.Setup(repo => repo.GetTasksAsync(query)).ReturnsAsync(tasks);

            var result = await _controller.GetTasks(query);

            var okResult = Assert.IsType<OkObjectResult>(result);
            var response = Assert.IsType<TaskSystemResponse<IEnumerable<TaskItem>>>(okResult.Value);
            Assert.Equal(ResponseMessages.GET_TASKS_SUCCESS, response.Message);
            Assert.Equal(tasks, response.Data);
        }

        [Fact]
        public async Task GetTasks_ReturnsNotFound_WhenNoTasksExist () {
            var query = new QueryObject();
            var emptyTasks = new List<TaskItem>();

            _mockRepo.Setup(repo => repo.GetTasksAsync(query)).ReturnsAsync(emptyTasks);

            var result = await _controller.GetTasks(query);

            var notFoundResult = Assert.IsType<NotFoundObjectResult>(result);
            var response = Assert.IsType<TaskSystemResponse<string>>(notFoundResult.Value);
            Assert.Equal(ResponseMessages.GET_TASKS_FAILED, response.Message);
            Assert.Null(response.Data);
        }

        [Fact]
        public async Task GetTaskById_ReturnsOkResult_WhenTaskExists()
        {
            int taskId = 1;
            var task = new TaskItem
            {
                Id = 1,
                TaskType = TaskItemType.Maintenance,
                Description = "There is a leak in the bathroom sink that needs urgent attention.",
                RoomId = 101,
                RequesterId = 1,
                AssigneeId = 2,
                Status = TaskItemStatus.InProgress,
                CreatedAt = new DateTime(2024, 10, 10, 10, 30, 0)
            };
            _mockRepo.Setup(repo => repo.GetTaskByIdAsync(taskId)).ReturnsAsync(task);

            var result = await _controller.GetTaskById(taskId);

            var okResult = Assert.IsType<OkObjectResult>(result);
            var response = Assert.IsType<TaskSystemResponse<TaskItem>>(okResult.Value);
            Assert.Equal(ResponseMessages.GET_TASK_SUCCESS, response.Message);
            Assert.Equal(task, response.Data);
        }

        [Fact]
        public async Task GetTaskById_ReturnsNotFound_WhenTaskDoesNotExist()
        {
            int taskId = 1;
            TaskItem? noTask = null;
            _mockRepo.Setup(repo => repo.GetTaskByIdAsync(taskId)).ReturnsAsync(noTask);

            var result = await _controller.GetTaskById(taskId);

            var notFoundResult = Assert.IsType<NotFoundObjectResult>(result);
            var response = Assert.IsType<TaskSystemResponse<int>>(notFoundResult.Value);
            Assert.Equal(ResponseMessages.GET_TASK_FAILED, response.Message);
            Assert.Equal(taskId, response.Data);
        }

        [Fact]
        public async Task AddTask_ReturnsCreatedAtAction_WhenTaskIsValid()
        {
            var taskDto = new AddTaskDto { TaskType = TaskItemType.RoomCleaning, Description = "Room Cleaning", RoomId = 101, RequesterId = 40 };
            var newTask = new TaskItem
            {
                Id = 1,
                TaskType = TaskItemType.RoomCleaning,
                Description = "Room Cleaning",
                RoomId = 101,
                RequesterId = 40,
                AssigneeId = null,
                Status = TaskItemStatus.Pending,
                CreatedAt = DateTime.Now
            };
            _mockRepo.Setup(repo => repo.AddTaskAsync(It.IsAny<TaskItem>())).ReturnsAsync(newTask);

            var result = await _controller.AddTask(taskDto);

            var createdAtActionResult = Assert.IsType<CreatedAtActionResult>(result);
            var response = Assert.IsType<TaskSystemResponse<TaskItem>>(createdAtActionResult.Value);
            Assert.Equal(ResponseMessages.CREATE_TASK_SUCCESS, response.Message);
            Assert.Equal(newTask, response.Data);
            Assert.Equal(nameof(TaskSystemController.GetTaskById), createdAtActionResult.ActionName);

            Assert.True(createdAtActionResult.RouteValues?.ContainsKey("id"));
            var routeId = Assert.IsType<int>(createdAtActionResult.RouteValues?["id"]);
            Assert.Equal(newTask.Id, routeId);
        }

        [Fact]
        public async Task UpdateTask_ReturnsOkResult_WhenUpdateIsSuccessful()
        {
            int taskId = 1;
            var taskDto = new UpdateTaskDto { TaskType = TaskItemType.RoomCleaning, Description = "Room Cleaning", AssigneeId = 30, Status = TaskItemStatus.InProgress };
            var updatedTask = new TaskItem
            {
                Id = taskId,
                TaskType = TaskItemType.RoomCleaning,
                Description = "Description",
                RoomId = 101,
                RequesterId = 1,
                AssigneeId = 30,
                Status = TaskItemStatus.InProgress,
                CreatedAt = new DateTime(2024, 10, 10, 10, 30, 0)
            };

            _mockRepo.Setup(repo => repo.UpdateTaskAsync(taskId, taskDto)).ReturnsAsync(updatedTask);

            var result = await _controller.UpdateTask(taskId, taskDto);

            var okResult = Assert.IsType<OkObjectResult>(result);
            var response = Assert.IsType<TaskSystemResponse<TaskItem>>(okResult.Value);
            Assert.Equal(ResponseMessages.UPDATE_TASK_SUCCESS, response.Message);
            Assert.Equal(updatedTask, response.Data);
        }

        [Fact]
        public async Task DeleteTask_ReturnsOkResult_WhenTaskExists()
        {
            int taskId = 1;
            var existingTask = new TaskItem
            {
                Id = taskId,
                TaskType = TaskItemType.RoomCleaning,
                Description = "Description",
                RoomId = 101,
                RequesterId = 1,
                AssigneeId = 30,
                Status = TaskItemStatus.InProgress,
                CreatedAt = new DateTime(2024, 10, 10, 10, 30, 0)
            };
            _mockRepo.Setup(repo => repo.GetTaskByIdAsync(taskId)).ReturnsAsync(existingTask);
            _mockRepo.Setup(repo => repo.DeleteTaskAsync(taskId)).ReturnsAsync(true);

            var result = await _controller.DeleteTask(taskId);

            var okResult = Assert.IsType<OkObjectResult>(result);
            var response = Assert.IsType<TaskSystemResponse<string>>(okResult.Value);
            Assert.Equal(ResponseMessages.DELETE_TASK_SUCCESS, response.Message);
            _mockRepo.Verify(repo => repo.DeleteTaskAsync(taskId), Times.Once);
        }

        [Fact]
        public async Task DeleteTask_ReturnsNotFound_WhenTaskDoesNotExist()
        {
            int taskId = 1;
            TaskItem? noTask = null;
            _mockRepo.Setup(repo => repo.GetTaskByIdAsync(taskId)).ReturnsAsync(noTask);

            var result = await _controller.DeleteTask(taskId);

            var notFoundResult = Assert.IsType<NotFoundObjectResult>(result);
            var response = Assert.IsType<TaskSystemResponse<int>>(notFoundResult.Value);
            Assert.Equal(ResponseMessages.GET_TASK_FAILED, response.Message);
            Assert.Equal(taskId, response.Data);
            _mockRepo.Verify(repo => repo.DeleteTaskAsync(taskId), Times.Never);
        }

        [Fact]
        public async Task GetTasks_ReturnsBadRequest_WhenModelStateIsInvalid()
        {
            _controller.ModelState.AddModelError("Error", "Sample Error");

            var result = await _controller.GetTasks(new QueryObject());

            Assert.IsType<BadRequestObjectResult>(result);
        }

        [Fact]
        public async Task GetTask_ReturnsBadRequest_WhenModelStateIsInvalid()
        {
            _controller.ModelState.AddModelError("Error", "Sample Error");

            var result = await _controller.GetTaskById(1);

            Assert.IsType<BadRequestObjectResult>(result);
        }

        [Fact]
        public async Task AddTask_ReturnsBadRequest_WhenModelStateIsInvalid()
        {
            _controller.ModelState.AddModelError("Error", "Sample Error");

            var result = await _controller.AddTask(new AddTaskDto());

            Assert.IsType<BadRequestObjectResult>(result);
        }

        [Fact]
        public async Task UpdateTask_ReturnsBadRequest_WhenModelStateIsInvalid()
        {
            _controller.ModelState.AddModelError("Error", "Sample error");

            var actionResult = await _controller.UpdateTask(1, new UpdateTaskDto());

            Assert.IsType<BadRequestObjectResult>(actionResult);
        }

        [Fact]
        public async Task DeleteTask_ReturnsBadRequest_WhenModelStateIsInvalid()
        {
            _controller.ModelState.AddModelError("Error", "Sample error");

            var actionResult = await _controller.DeleteTask(1);

            Assert.IsType<BadRequestObjectResult>(actionResult);
        }
    }
}
