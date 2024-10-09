public static class Services
{
    private static IAmenityPersistence _amenityPersistence = null;

    public static IAmenityPersistence GetAmenityPersistence()
    {
        if(_amenityPersistence == null){
            _amenityPersistence = new StubAmenityPersistence();
        }

        return _amenityPersistence;
    }

    public static void clear(){
        _amenityPersistence = null;
    }
}