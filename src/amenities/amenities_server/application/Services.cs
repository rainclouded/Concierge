public static class Services
{
    private static IAmenityPersistence? _amenityPersistence = null;

    public static IAmenityPersistence GetAmenityPersistence()
    {
        _amenityPersistence ??= ConstructAmenityPersistence();
        return _amenityPersistence;
    }

    public static void clear()
    {
        _amenityPersistence = null;
    }

    private static IAmenityPersistence ConstructAmenityPersistence()
    {
        string? dbImplementation = Environment.GetEnvironmentVariable("DB_IMPLEMENTATION") ?? String.Empty;
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

        return amenityPersistence ?? new StubAmenityPersistence();
    }

    private static string PostgresConnectionString()
    {
        string? dbHost = Environment.GetEnvironmentVariable("DB_HOST") ?? String.Empty;
        string? dbPort = Environment.GetEnvironmentVariable("DB_PORT") ?? String.Empty;
        string? dbUsername = Environment.GetEnvironmentVariable("DB_USERNAME") ?? String.Empty;
        string? dbPassword = Environment.GetEnvironmentVariable("DB_PASSWORD") ?? String.Empty;
        return $"Host={dbHost}; Port={dbPort}; Username={dbUsername}; Password={dbPassword}";
    }
}
