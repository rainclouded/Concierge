public static class Services
{
    private static IAmenityPersistence _amenityPersistence = null;

    public static IAmenityPersistence GetAmenityPersistence()
    {
        _amenityPersistence??=ConfigureDB();
        return _amenityPersistence;
    }

    public static void clear(){
        _amenityPersistence = null;
    }
    
    private static IAmenityPersistence ConfigureDB(){
      string? dbImplentation = Environment.GetEnvironmentVariable("DB_IMPLEMENTATION");
      return ConfigureDB(dbImplentation ?? "");
    }
    
    private static IAmenityPersistence ConfigureDB(string dbImplentation)
    {
      IAmenityPersistence? amenityPersistence = null;
      if (dbImplentation == "POSTGRES")
      { 
         Console.WriteLine("Attempting to connect to Postgres");
         string? dbHost = Environment.GetEnvironmentVariable("DB_HOST") ?? "";
         string? dbPort = Environment.GetEnvironmentVariable("DB_PORT") ?? "";
         string? dbUsername = Environment.GetEnvironmentVariable("DB_USERNAME") ?? "";
         string? dbPassword = Environment.GetEnvironmentVariable("DB_PASSWORD") ?? "";
         try
         {
            amenityPersistence = new PostgresAmenityPersistence($"Host={dbHost}; Port={dbPort}; Username={dbUsername}; Password={dbPassword}");
            Console.WriteLine("Postgress Connected successfully");
         }
         catch (InvalidOperationException)
         {
            amenityPersistence = null; //use stub if connection failed
            Console.WriteLine("Postgress Failed");
         }
      }
      
      return amenityPersistence ?? new StubAmenityPersistence();
    }
}
