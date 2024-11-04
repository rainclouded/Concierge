using System.ComponentModel.DataAnnotations;

namespace task_system_server.Dtos;

public class UpdateAssigneeDto
{
    [Required]
    public int AssigneeId { get; set; }
}
