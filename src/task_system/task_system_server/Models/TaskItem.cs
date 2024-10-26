using System.ComponentModel.DataAnnotations;

namespace task_system_server.Models;

public class TaskItem
{
	public int Id { get; set; }
	public string Title { get; set; }  = string.Empty;
	public string Description { get; set; } = string.Empty;
	public int RoomId { get; set; }
	public int RequesterId { get; set; }
	public int? AssigneeId { get; set; }
	public string Status { get; set; } = "Pending";
	public DateTime CreatedAt { get; set; } = DateTime.Now;

	public TaskItem() { }
}
