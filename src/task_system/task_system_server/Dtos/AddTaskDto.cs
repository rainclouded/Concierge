using System;
using System.ComponentModel.DataAnnotations;
using task_system_server.Models;

namespace task_system_server.Dtos;

public class AddTaskDto
{
    [Required]
	public TaskItemType TaskType { get; set; }

    [Required]
    [MaxLength(400, ErrorMessage = "Description cannot be over 400 over characters")]
	public string Description { get; set; } = string.Empty;

    [Required]
    [Range(100, 500)]
    public int RoomId { get; set; }
    
    [Required]
    [Range(1, 500)]
	public int RequesterId { get; set; }
}
