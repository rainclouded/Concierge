namespace task_system_server.Models;

public enum TaskItemType
{
	RoomCleaning,
	Maintenance,
	FoodDelivery,
	WakeUpCall,
	LaundryService,
	SpaAndMassage
}

public enum TaskItemStatus
{
	Pending,
	InProgress,
	Completed
}

public class TaskItem
{
	public int Id { get; set; }
	public TaskItemType? TaskType { get; set; }
	public string Description { get; set; } = string.Empty;
	public int RoomId { get; set; }
	public int RequesterId { get; set; }
	public int? AssigneeId { get; set; }
	public TaskItemStatus Status { get; set; } = TaskItemStatus.Pending;
	public DateTime CreatedAt { get; set; } = DateTime.Now;

	public TaskItem() { }
}
