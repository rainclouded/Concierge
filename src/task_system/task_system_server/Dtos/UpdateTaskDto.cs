using System;
using System.ComponentModel.DataAnnotations;

namespace task_system_server.Dtos;

public class UpdateTaskDto
{
    [Required]
	public string Title { get; set; } = string.Empty;

    [Required]
    [MaxLength(400, ErrorMessage = "Description cannot be over 400 over characters")]
	public string Description { get; set; } = string.Empty;

    [Required]
    [Range(1, 500, ErrorMessage = "There are only 500 rooms")]
	public int AssigneeId { get; set; }

    [Required]
    [RegularExpression("^(Pending|Complete|In Progress)$", ErrorMessage = "Status must be either 'Pending', 'Complete', or 'In Progress'.")]
	public string Status { get; set; } = string.Empty;
}
