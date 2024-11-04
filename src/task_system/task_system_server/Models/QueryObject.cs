namespace task_system_server.Models;

public class QueryObject
{
    public int? RoomId { get; set; } = null;
    public int? RequesterId { get; set; } = null;
    public int? AssigneeId { get; set; } = null;
    public TaskItemStatus? Status { get; set; } = null;
    public int? Year { get; set; } = null;
    public int? Month { get; set; } = null;
    public int? Day { get; set; } = null;
    public bool SortAscending { get; set; } = true;
}
