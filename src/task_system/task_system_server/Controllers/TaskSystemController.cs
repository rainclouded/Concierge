using Microsoft.AspNetCore.Mvc;
using task_system_server.Models;
using task_system_server.Interfaces;
using task_system_server.Dtos;
using task_system_server.Mappers;
using Microsoft.AspNetCore.Cors;
using task_system_server.Validators;

namespace task_system_server.Controllers
{
    [Route("tasks")]
    [ApiController]
    [EnableCors("AllowAll")]
    public class TaskSystemController : ControllerBase
    {
        private readonly ITaskSystemRepository _taskSystemRepository;
        private readonly IPermissionValidator _permissionValidator;

        /* 
        Constructor for TaskSystemController.
        Initializes the controller with task system repository and permission validator.

        Args:
        taskSystemRepository: Repository for interacting with tasks in the system.
        permissionValidator: Validator for checking if the API key has appropriate permissions.
        */
        public TaskSystemController(ITaskSystemRepository taskSystemRepository, IPermissionValidator permissionValidator)
        {
            _taskSystemRepository = taskSystemRepository;
            _permissionValidator = permissionValidator;
        }

        /* 
        GetTasks retrieves tasks based on the provided query parameters.
        It validates the API key and permission to view tasks.

        Args:
        query: Contains filter parameters for the tasks, such as AccountId.

        Returns:
        A response with the list of tasks or an error message.
        */
        [HttpGet]
        public async Task<IActionResult> GetTasks([FromQuery] QueryObject query)
        {
            if (!Request.Headers.TryGetValue("X-API-Key", out var apiKey))
                return Unauthorized(new TaskSystemResponse<string>(ResponseMessages.UNAUTHORIZED, null));

            if (!_permissionValidator.ValidatePermissions(PermissionNames.VIEW_TASKS, apiKey!, out var sessionData))
            {
                if(sessionData == null)
                    return StatusCode(StatusCodes.Status500InternalServerError, "An unexpected error occurred.");
                
                // if client asks for a different account's task details, return unauthorized
                if(query.AccountId.HasValue)
                    return Unauthorized(new TaskSystemResponse<string>(ResponseMessages.UNAUTHORIZED, null));

                query.AccountId = sessionData.AccountId;
            }

            if (!ModelState.IsValid)
                return BadRequest(ModelState);
                
            var tasks = await _taskSystemRepository.GetTasksAsync(query);

            if (!tasks.Any())
                return NotFound(new TaskSystemResponse<string>(ResponseMessages.GET_TASKS_FAILED, null));

            return Ok(new TaskSystemResponse<IEnumerable<TaskItem>>(ResponseMessages.GET_TASKS_SUCCESS, tasks));
        }

        /* 
        GetTaskById retrieves a specific task by its ID.

        Args:
        taskId: The ID of the task to retrieve.

        Returns:
        A response with the task details or an error message if not found.
        */
        [HttpGet("{id}")]
        public async Task<IActionResult> GetTaskById([FromRoute] int taskId)
        {
            if (!Request.Headers.TryGetValue("X-API-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.VIEW_TASKS, apiKey!, out var sessionData))
                return Unauthorized(new TaskSystemResponse<string>(ResponseMessages.UNAUTHORIZED, null));

            if (!ModelState.IsValid)
                return BadRequest(ModelState);

            var task = await _taskSystemRepository.GetTaskByIdAsync(taskId);

            if (task == null)
            {
                return NotFound(new TaskSystemResponse<int>(ResponseMessages.GET_TASK_FAILED, taskId));
            }

            return Ok(new TaskSystemResponse<TaskItem>(ResponseMessages.GET_TASK_SUCCESS, task));
        }

        /* 
        AddTask creates a new task in the system.

        Args:
        taskDto: Data transfer object containing the details for the new task.

        Returns:
        A response with the created task or an error message if creation fails.
        */
        [HttpPost]
        public async Task<IActionResult> AddTask([FromBody] AddTaskDto taskDto)
        {
            
            if (!Request.Headers.TryGetValue("X-API-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.CREATE_TASKS, apiKey!, out var sessionData))
                return Unauthorized(new TaskSystemResponse<string>(ResponseMessages.UNAUTHORIZED, null));

            if (!ModelState.IsValid)
                return BadRequest(ModelState);

            var newTask = taskDto.ToTaskFromAddTaskDto();

            var createdTask = await _taskSystemRepository.AddTaskAsync(newTask);

            return CreatedAtAction(
                nameof(GetTaskById),
                new { id = createdTask.Id },
                new TaskSystemResponse<TaskItem>(ResponseMessages.CREATE_TASK_SUCCESS, createdTask)
            );
        }

        /* 
        UpdateTask updates the details of an existing task.

        Args:
        id: The ID of the task to be updated.
        taskDto: Data transfer object containing updated task details.

        Returns:
        A response with the updated task or an error message if update fails.
        */
        [HttpPut("{id}")]
        public async Task<IActionResult> UpdateTask([FromRoute] int id, [FromBody] UpdateTaskDto taskDto)
        {
            if (!Request.Headers.TryGetValue("X-API-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.EDIT_TASKS, apiKey!, out var sessionData))
                return Unauthorized(new TaskSystemResponse<string>(ResponseMessages.UNAUTHORIZED, null));

            if (!ModelState.IsValid)
                return BadRequest(ModelState);

            var updatedTask = await _taskSystemRepository.UpdateTaskAsync(id, taskDto);

            return Ok(new TaskSystemResponse<TaskItem>(ResponseMessages.UPDATE_TASK_SUCCESS, updatedTask));
        }

        /* 
        UpdateAssignee changes the assignee of a task.

        Args:
        id: The ID of the task whose assignee is to be updated.
        assigneeDto: Data transfer object containing the new assignee ID.

        Returns:
        A response with the updated task or an error message if update fails.
        */
        [HttpPatch("{id}/assignee")]
        public async Task<IActionResult> UpdateAssignee([FromRoute] int id, [FromBody] UpdateAssigneeDto assigneeDto)
        {
            if (!Request.Headers.TryGetValue("X-API-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.EDIT_TASKS, apiKey!, out var sessionData))
                return Unauthorized(new TaskSystemResponse<string>(ResponseMessages.UNAUTHORIZED, null));
                
            if (!ModelState.IsValid)
                return BadRequest(ModelState);

            var updatedTask = await _taskSystemRepository.UpdateAssigneeAsync(id, assigneeDto.AssigneeId);

            return Ok(new TaskSystemResponse<TaskItem>(ResponseMessages.UPDATE_TASK_SUCCESS, updatedTask));
        }

        /* 
        DeleteTask removes a task from the system by its ID.

        Args:
        id: The ID of the task to be deleted.

        Returns:
        A response indicating whether the task was deleted successfully or an error message if deletion fails.
        */
        [HttpDelete("{id}")]
        public async Task<IActionResult> DeleteTask([FromRoute] int id)
        {
            if (!Request.Headers.TryGetValue("X-API-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.DELETE_TASKS, apiKey!, out var sessionData))
                return Unauthorized(new TaskSystemResponse<string>(ResponseMessages.UNAUTHORIZED, null));

            if (!ModelState.IsValid)
                return BadRequest(ModelState);

            if (await _taskSystemRepository.GetTaskByIdAsync(id) == null)
            {
                return NotFound(new TaskSystemResponse<int>(ResponseMessages.GET_TASK_FAILED, id));
            }

            await _taskSystemRepository.DeleteTaskAsync(id);
            
            return Ok(new TaskSystemResponse<string>(ResponseMessages.DELETE_TASK_SUCCESS, null));
        }
    }
}
