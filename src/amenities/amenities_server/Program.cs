using amenities_server.services;
using amenities_server.validators;

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
builder.Services.AddHttpClient();
builder.Services.AddSingleton<PermissionClient>();
builder.Services.AddSingleton<IPermissionValidator, PermissionValidator>();

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
public partial class Program { } //For running integration tests