using task_system_server.Dtos;
using task_system_server.Models;

namespace task_system_server.Mappers;

public static class TaskMapper
{
    public static TaskItem ToTaskFromAddTaskDto (this AddTaskDto taskDto) {
        return new TaskItem {
            TaskType = taskDto.TaskType,
            Description = taskDto.Description,
            RoomId = taskDto.RoomId,
            RequesterId = taskDto.RequesterId,
        };
    }
}
