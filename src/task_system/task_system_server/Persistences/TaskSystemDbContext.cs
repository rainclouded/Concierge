using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion;
using task_system_server.Models;

namespace task_system_server.Persistences;

public class TaskSystemDbContext : DbContext
{
    /*
    Constructor for TaskSystemDbContext
    Args:
      DbContextOptions<TaskSystemDbContext> dbContextOptions: Options to configure the context with
    Returns:
      None
    */
    public TaskSystemDbContext(DbContextOptions<TaskSystemDbContext> dbContextOptions)
    : base(dbContextOptions)
    {   
    }

    /*
    DbSet property for accessing TaskItems
    Args:
      None
    Returns:
      DbSet<TaskItem>: A collection of TaskItem entities in the database
    */
    public DbSet<TaskItem> Tasks { get; set; }

    /*
    Configures the model and applies any required configuration settings
    Args:
      ModelBuilder modelBuilder: The builder used to configure the model
    Returns:
      None
    */
    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        base.OnModelCreating(modelBuilder);

        // Apply the HasColumnType configuration to all DateTime properties
        foreach (var entityType in modelBuilder.Model.GetEntityTypes())
        {
            foreach (var property in entityType.GetProperties())
            {
                // Check if the property type is DateTime
                if (property.ClrType == typeof(DateTime))
                {
                    /*
                    Sets the value converter for DateTime properties to convert between 
                    local time and UTC time when saving and reading from the database
                    Args:
                      v: DateTime value to be converted to UTC
                      v: DateTime value to be converted back to local time
                    Returns:
                      None
                    */
                    property.SetValueConverter(new ValueConverter<DateTime, DateTime>(
                        v => v.ToUniversalTime(), // Convert Local to UTC on save
                        v => DateTime.SpecifyKind(v, DateTimeKind.Local) // Convert back to Local on read
                    ));
                }
            }
        }
    }
}
