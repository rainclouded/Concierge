public static class Services
{
    private static IAmenityPersistence _amenityPersistence = null;

    public static IAmenityPersistence GetAmenityPersistence(string forProduction)
    {
        if(_amenityPersistence == null){
            if(forProduction.Equals("test"))
            {
                _amenityPersistence = new StubAmenityPersistence();
            }
        }
        return _amenityPersistence;
    }
}