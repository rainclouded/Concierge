using Microsoft.EntityFrameworkCore;
using task_system_server.Interfaces;
using task_system_server.Persistences;
using task_system_server.Repositories;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddControllers();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddCors(options =>
{
    options.AddPolicy("AllowAll",
        policy=>
        {
            policy.AllowAnyOrigin()      // Allow requests from any origin
                   .AllowAnyMethod()      // Allow any HTTP method (GET, POST, PUT, DELETE, etc.)
                   .AllowAnyHeader();     // Allow any headers in the request
        });
});

builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();


var dbImplementation = Environment.GetEnvironmentVariable("DB_IMPLEMENTATION");

if(dbImplementation == "POSTGRES"){
    string? dbHost = Environment.GetEnvironmentVariable("DB_HOST") ?? string.Empty;
    string? dbPort = Environment.GetEnvironmentVariable("DB_PORT") ?? string.Empty;
    string? dbUsername = Environment.GetEnvironmentVariable("DB_USERNAME") ?? string.Empty;
    string? dbPassword = Environment.GetEnvironmentVariable("DB_PASSWORD") ?? string.Empty;

    string connectionString = $"Host={dbHost}; Port={dbPort}; Username={dbUsername}; Password={dbPassword}";
    
    // New instance per request - fresh database context
    builder.Services.AddDbContext<TaskSystemDbContext>(options =>
        options.UseNpgsql(connectionString));
    builder.Services.AddScoped<ITaskSystemRepository, PostgresTaskSystemRepository>();
} else {
    // One instance for all requests - maintains in-memory state
    builder.Services.AddSingleton<ITaskSystemRepository, StubTaskSystemRepository>();
}

/*
// Configure repository. Change to true/false in appsetting.json
bool useStubRepository = builder.Configuration.GetValue<bool>("InDevelopment");

if (useStubRepository){
    // One instance for all requests - maintains in-memory state
    builder.Services.AddSingleton<ITaskSystemRepository, StubTaskSystemRepository>();
} else {
    // New instance per request - fresh database context
    builder.Services.AddDbContext<TaskSystemDbContext>(options =>
        options.UseNpgsql(builder.Configuration.GetConnectionString("DefaultConnection")));
    builder.Services.AddScoped<ITaskSystemRepository, PostgresTaskSystemRepository>();
}
*/

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseRouting(); 
app.UseCors("AllowAll");

app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();