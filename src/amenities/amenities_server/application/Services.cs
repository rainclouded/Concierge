using amenities_server.persistence;
using amenities_server.services;
using amenities_server.validators;
using Microsoft.Extensions.DependencyInjection;

namespace amenities_server.application
{
    public static class Services
    {
        private static IAmenityPersistence? _amenityPersistence = null;
        private static IPermissionValidator? _permissionValidator = null;

        public static IAmenityPersistence GetAmenityPersistence()
        {
            //return a singleton of the amenity persistence, construct one if it does not exist
            _amenityPersistence ??= ConstructAmenityPersistence();
            return _amenityPersistence;
        }

        public static IPermissionValidator GetPermissionValidator(IHttpClientFactory httpFactory)
        {
            //return a singleton of the permission validator, construct one if it does not exist
            _permissionValidator ??= ConstructPermissionValidator(httpFactory);
            return _permissionValidator;
        }

        public static void SetAmenityPersistence(IAmenityPersistence amenityPersistence)
        {
            _amenityPersistence = amenityPersistence;
        }

        public static void Clear()
        {
            _amenityPersistence = null;
        }

        private static IAmenityPersistence ConstructAmenityPersistence()
        {
            //construct a instance of the amenity persistence, creates a mock or postgres implementation based on the env variable 'DB_IMPLEMENTATION' stated in the yaml file
            string? dbImplementation = Environment.GetEnvironmentVariable("DB_IMPLEMENTATION") ?? string.Empty;
            IAmenityPersistence? amenityPersistence = null;
            if (dbImplementation == "POSTGRES")
            {
                Console.WriteLine("Attempting to connect to Postgres");
                try
                {
                    amenityPersistence = new PostgresAmenityPersistence(PostgresConnectionString());
                    Console.WriteLine("Postgress Connected successfully");
                }
                catch (InvalidOperationException)
                {
                    amenityPersistence = null;
                    Console.WriteLine("Postgress Failed");
                }
            }
            else if (dbImplementation == "MOCK")
            {
                amenityPersistence = new StubAmenityPersistence();
            }

            return amenityPersistence ?? new StubAmenityPersistence();
        }

        private static IPermissionValidator ConstructPermissionValidator(IHttpClientFactory httpFactory)
        {
            Console.WriteLine("Constructing PermissionValidator");
            var PermCli = new PermissionClient(httpFactory);
            return new PermissionValidator(PermCli);
        }

        private static string PostgresConnectionString()
        {
            //if a postgres connection is to be made, grab connection details from the environment as stated in the yaml file
            string? dbHost = Environment.GetEnvironmentVariable("DB_HOST") ?? string.Empty;
            string? dbPort = Environment.GetEnvironmentVariable("DB_PORT") ?? string.Empty;
            string? dbUsername = Environment.GetEnvironmentVariable("DB_USERNAME") ?? string.Empty;
            string? dbPassword = Environment.GetEnvironmentVariable("DB_PASSWORD") ?? string.Empty;
            return $"Host={dbHost}; Port={dbPort}; Username={dbUsername}; Password={dbPassword}";
        }
    }
}