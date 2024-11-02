using amenities_server.persistence;
using amenities_server.validators;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Mvc.Testing;
using Microsoft.Extensions.DependencyInjection;

namespace amenities_db_integration_test
{
    internal class WebAppFactoryForTests(string connectionString) : WebApplicationFactory<Program>
    {
        private readonly string _connectionString = connectionString;

        protected override void ConfigureWebHost(IWebHostBuilder builder)
        {
            builder.ConfigureServices(services =>
            {
                amenities_server.application.Services.SetAmenityPersistence(new PostgresAmenityPersistence(_connectionString));
                services.AddSingleton<IPermissionValidator, MockPermissionValidator>(); //Set the service provider to null for permissionValidation
            });
        }
    }
}
