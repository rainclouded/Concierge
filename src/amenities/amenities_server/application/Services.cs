using amenities_server.persistence;

namespace amenities_server.application
{
    public static class Services
    {
        private static IAmenityPersistence? _amenityPersistence = null;

        public static IAmenityPersistence GetAmenityPersistence()
        {
            _amenityPersistence ??= ConstructAmenityPersistence();
            return _amenityPersistence;
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

        private static string PostgresConnectionString()
        {
            string? dbHost = Environment.GetEnvironmentVariable("DB_HOST") ?? string.Empty;
            string? dbPort = Environment.GetEnvironmentVariable("DB_PORT") ?? string.Empty;
            string? dbUsername = Environment.GetEnvironmentVariable("DB_USERNAME") ?? string.Empty;
            string? dbPassword = Environment.GetEnvironmentVariable("DB_PASSWORD") ?? string.Empty;
            return $"Host={dbHost}; Port={dbPort}; Username={dbUsername}; Password={dbPassword}";
        }
    }
}