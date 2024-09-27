public static class AmenityValidator
{
    private static IAmenityPersistence amenityPersistence;
    
    public static bool ValidateAmenityParameters(Amenity amenity)
    {
        if (amenity == null)
        {
            return false;
        }

        if(amenity.AmenityID < 0){
            return false;
        }

        if(string.IsNullOrWhiteSpace(amenity.AmenityName)){
            return false;
        }

        if(string.IsNullOrWhiteSpace(amenity.AmenityDescription)){
            return false;
        }

        if(amenity.StartTime >= amenity.EndTime){
            return false;
        }

        return true;;
    }

    public static bool ValidateNewAmenity(Amenity amenity)
    {
        //get recent instance of persistence
        amenityPersistence = Services.GetAmenityPersistence();

        if (!ValidateAmenityParameters(amenity)){
            return false;
        }

        return amenityPersistence.GetAmenityByID(amenity.AmenityID) == null;
    }

    public static bool ValidateExistingAmenity(Amenity amenity)
    {
        //get recent instance of persistence
        amenityPersistence = Services.GetAmenityPersistence();

        if (!ValidateAmenityParameters(amenity)){
            return false;
        }

        return amenityPersistence.GetAmenityByID(amenity.AmenityID) != null;
    }
}