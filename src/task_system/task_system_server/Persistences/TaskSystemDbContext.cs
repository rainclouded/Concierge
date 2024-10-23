using Microsoft.EntityFrameworkCore;
using task_system_server.Models;

namespace task_system_server.Persistences;

public class TaskSystemDbContext : DbContext
{
    public TaskSystemDbContext(DbContextOptions<TaskSystemDbContext> dbContextOptions)
    : base(dbContextOptions)
    {   
    }

    public DbSet<TaskItem> Tasks { get; set; }
}
