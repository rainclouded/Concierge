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

        /*
        Returns a singleton of the amenity persistence, constructs one if it does not exist
        Args:
            None
        Returns:
            IAmenityPersistence: The singleton instance of AmenityPersistence
        */
        public static IAmenityPersistence GetAmenityPersistence()
        {
            //return a singleton of the amenity persistence, construct one if it does not exist
            _amenityPersistence ??= ConstructAmenityPersistence();
            return _amenityPersistence;
        }

        /*
        Returns a singleton of the permission validator, constructs one if it does not exist
        Args:
            IHttpClientFactory httpFactory: A factory to create HTTP clients for permission validation
        Returns:
            IPermissionValidator: The singleton instance of PermissionValidator
        */
        public static IPermissionValidator GetPermissionValidator(IHttpClientFactory httpFactory)
        {
            //return a singleton of the permission validator, construct one if it does not exist
            _permissionValidator ??= ConstructPermissionValidator(httpFactory);
            return _permissionValidator;
        }

        /*
        Sets the AmenityPersistence instance
        Args:
            IAmenityPersistence amenityPersistence: The instance of AmenityPersistence to set
        Returns:
            void
        */
        public static void SetAmenityPersistence(IAmenityPersistence amenityPersistence)
        {
            _amenityPersistence = amenityPersistence;
        }

        /*
        Clears the stored AmenityPersistence instance
        Args:
            None
        Returns:
            void
        */
        public static void Clear()
        {
            _amenityPersistence = null;
        }

        /*
        Constructs the AmenityPersistence instance based on environment variables
        Args:
            None
        Returns:
            IAmenityPersistence: The constructed instance of AmenityPersistence
        */
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

        /*
        Constructs the PermissionValidator instance
        Args:
            IHttpClientFactory httpFactory: A factory to create HTTP clients for permission validation
        Returns:
            IPermissionValidator: The constructed instance of PermissionValidator
        */
        private static IPermissionValidator ConstructPermissionValidator(IHttpClientFactory httpFactory)
        {
            Console.WriteLine("Constructing PermissionValidator");
            var PermCli = new PermissionClient(httpFactory);
            return new PermissionValidator(PermCli);
        }

        /*
        Returns the Postgres connection string using environment variables
        Args:
            None
        Returns:
            string: The constructed Postgres connection string
        */
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