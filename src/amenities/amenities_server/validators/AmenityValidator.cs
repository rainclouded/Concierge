using amenities_server.application;
using amenities_server.model;
using amenities_server.persistence;

namespace amenities_server.validators
{
    public static class AmenityValidator
    {
        private static IAmenityPersistence amenityPersistence;

        public static bool ValidateAmenityParameters(Amenity amenity)
        {
            if (amenity == null)
            {
                return false;
            }

            if (string.IsNullOrWhiteSpace(amenity.Name))
            {
                return false;
            }

            if (string.IsNullOrWhiteSpace(amenity.Description))
            {
                return false;
            }

            return true;
        }

        public static bool ValidateNewAmenity(Amenity amenity)
        {
            //get recent instance of persistence
            amenityPersistence = Services.GetAmenityPersistence();

            if (!ValidateAmenityParameters(amenity))
            {
                return false;
            }

            return amenity.Id < 0 || amenityPersistence.GetAmenityByID(amenity.Id) == null;
        }
    }
}