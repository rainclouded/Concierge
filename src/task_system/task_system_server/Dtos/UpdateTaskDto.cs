using System.ComponentModel.DataAnnotations;
using task_system_server.Models;

namespace task_system_server.Dtos;

public class UpdateTaskDto
{
    [Required]
	public TaskItemType TaskType { get; set; }

    [Required]
    [MaxLength(400, ErrorMessage = "Description cannot be over 400 over characters")]
	public string Description { get; set; } = string.Empty;

    [Required]
    [Range(1, 500, ErrorMessage = "There are only 500 rooms")]
	public int AssigneeId { get; set; }

    [Required]
	public TaskItemStatus Status { get; set; }
}
