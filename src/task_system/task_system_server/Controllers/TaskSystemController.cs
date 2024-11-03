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

        public TaskSystemController(ITaskSystemRepository taskSystemRepository, IPermissionValidator permissionValidator)
        {
            _taskSystemRepository = taskSystemRepository;
            _permissionValidator = permissionValidator;
        }

        //GET: /tasks
        [HttpGet]
        public async Task<IActionResult> GetTasks([FromQuery] QueryObject query)
        {
            if (!Request.Headers.TryGetValue("X-API-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.VIEW_TASKS, apiKey!))
                return Unauthorized(new TaskSystemResponse<string>(ResponseMessages.UNAUTHORIZED, null));

            if (!ModelState.IsValid)
                return BadRequest(ModelState);

            var tasks = await _taskSystemRepository.GetTasksAsync(query);

            if (!tasks.Any())
            {
                return NotFound(new TaskSystemResponse<string>(ResponseMessages.GET_TASKS_FAILED, null));
            }

            return Ok(new TaskSystemResponse<IEnumerable<TaskItem>>(ResponseMessages.GET_TASKS_SUCCESS, tasks));
        }

        //GET: /tasks/{id}
        [HttpGet("{id}")]
        public async Task<IActionResult> GetTaskById([FromRoute] int id)
        {
            if (!Request.Headers.TryGetValue("X-API-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.VIEW_TASKS, apiKey!))
                return Unauthorized(new TaskSystemResponse<string>(ResponseMessages.UNAUTHORIZED, null));

            if (!ModelState.IsValid)
                return BadRequest(ModelState);

            var task = await _taskSystemRepository.GetTaskByIdAsync(id);

            if (task == null)
            {
                return NotFound(new TaskSystemResponse<int>(ResponseMessages.GET_TASK_FAILED, id));
            }

            return Ok(new TaskSystemResponse<TaskItem>(ResponseMessages.GET_TASK_SUCCESS, task));
        }

        //POST: /tasks
        [HttpPost]
        public async Task<IActionResult> AddTask([FromBody] AddTaskDto taskDto)
        {
            
            if (!Request.Headers.TryGetValue("X-API-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.CREATE_TASKS, apiKey!))
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

        //PUT: /tasks/{id}
        [HttpPut("{id}")]
        public async Task<IActionResult> UpdateTask([FromRoute] int id, [FromBody] UpdateTaskDto taskDto)
        {
            if (!Request.Headers.TryGetValue("X-API-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.EDIT_TASKS, apiKey!))
                return Unauthorized(new TaskSystemResponse<string>(ResponseMessages.UNAUTHORIZED, null));

            if (!ModelState.IsValid)
                return BadRequest(ModelState);

            var updatedTask = await _taskSystemRepository.UpdateTaskAsync(id, taskDto);

            return Ok(new TaskSystemResponse<TaskItem>(ResponseMessages.UPDATE_TASK_SUCCESS, updatedTask));
        }

        //DELETE: /tasks/{id}
        [HttpDelete("{id}")]
        public async Task<IActionResult> DeleteTask([FromRoute] int id)
        {
            if (!Request.Headers.TryGetValue("X-API-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.DELETE_TASKS, apiKey!))
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
