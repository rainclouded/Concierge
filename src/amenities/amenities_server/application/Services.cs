public static class Services
{
    private static IAmenityPersistence _amenityPersistence = null;
    private static IPermissionValidator _permissionValidator = null;

    public static IAmenityPersistence GetAmenityPersistence()
    {
        if(_amenityPersistence == null){
            _amenityPersistence = new StubAmenityPersistence();
        }

        return _amenityPersistence;
    }

    public static IPermissionValidator GetPermissionValidator()
    {
        if(_permissionValidator == null){
            _permissionValidator = new PermissionValidator();
        }

        return _permissionValidator;
    }

    public static void clear(){
        _amenityPersistence = null;
        _permissionValidator = null;
    }
}