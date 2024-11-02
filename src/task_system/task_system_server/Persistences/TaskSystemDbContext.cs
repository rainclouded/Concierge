using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion;
using task_system_server.Models;

namespace task_system_server.Persistences;

public class TaskSystemDbContext : DbContext
{
    public TaskSystemDbContext(DbContextOptions<TaskSystemDbContext> dbContextOptions)
    : base(dbContextOptions)
    {   
    }

    public DbSet<TaskItem> Tasks { get; set; }

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        base.OnModelCreating(modelBuilder);

        // Apply the HasColumnType configuration to all DateTime properties
        foreach (var entityType in modelBuilder.Model.GetEntityTypes())
        {
            foreach (var property in entityType.GetProperties())
            {
                if (property.ClrType == typeof(DateTime))
                {
                    property.SetValueConverter(new ValueConverter<DateTime, DateTime>(
                        v => v.ToUniversalTime(), // Convert Local to UTC on save
                        v => DateTime.SpecifyKind(v, DateTimeKind.Local) // Convert back to Local on read
                    ));
                }
            }
        }
    }
}
