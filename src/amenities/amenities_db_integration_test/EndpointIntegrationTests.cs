using Npgsql;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace amenities_db_integration_test
{
    internal class EndpointIntegrationTests
    {
        private PostgresAmenityPersistence? _persistence;
        private string? _connectionString;

        [OneTimeSetUp]
        public void OneTimeSetUp()
        {
            //string dbHost = Environment.GetEnvironmentVariable("DB_HOST") ?? String.Empty;
            //string dbPort = Environment.GetEnvironmentVariable("DB_PORT") ?? String.Empty;
            //string dbUsername = Environment.GetEnvironmentVariable("DB_USERNAME") ?? String.Empty;
            //string dbPassword = Environment.GetEnvironmentVariable("DB_PASSWORD") ?? String.Empty;
            //string connnectionString = $"Host={dbHost}; Port={dbPort}; Username={dbUsername}; Password={dbPassword}";
            _connectionString = "Host=127.0.0.1; Port=50014; Username=postgres; Password=sa";

            _persistence = new PostgresAmenityPersistence(_connectionString);
            Services.SetAmenityPersistence(_persistence);
        }

        [SetUp]
        public void Setup()
        {
            using var connection = new NpgsqlConnection("Host=127.0.0.1; Port=50014; Username=postgres; Password=sa");
            connection.Open();
            using var command = new NpgsqlCommand("delete from amenities", connection);
            command.ExecuteReader();
        }


    }
}
