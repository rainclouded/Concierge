public static class Services
{
    public static IAmenityPersistence GetAmenityPersistence(bool forProduction)
    {
        if (forProduction)
        {
            //return mongoDB
        }
        else
        {
            return new StubAmenityPersistence();
        }
    }
}