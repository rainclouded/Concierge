using task_system_server.Dtos;
using task_system_server.Models;

namespace task_system_server.Mappers
{
    /*
    The TaskMapper class contains mapping methods between data transfer objects (DTOs) and models.
    It provides functionality to convert an AddTaskDto to a TaskItem object.

    */
    public static class TaskMapper
    {
        /*
        Converts an AddTaskDto to a TaskItem model.

        This method takes an AddTaskDto object and maps its properties to a new TaskItem object.
        The TaskItem is initialized with the TaskType, Description, RoomId, and RequesterId
        from the AddTaskDto, and the CreatedAt property is set to the current UTC time.
        The Status is set to Pending by default.

        Args:
        taskDto: The AddTaskDto object containing task details.

        Returns:
        A new TaskItem object with the values from the AddTaskDto.
        */
        public static TaskItem ToTaskFromAddTaskDto(this AddTaskDto taskDto)
        {
            return new TaskItem
            {
                TaskType = taskDto.TaskType,
                Description = taskDto.Description,
                RoomId = taskDto.RoomId,
                RequesterId = taskDto.RequesterId,

                CreatedAt = DateTime.UtcNow,
                Status = TaskItemStatus.Pending
            };
        }
    }
}
