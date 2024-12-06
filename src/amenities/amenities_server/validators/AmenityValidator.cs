using amenities_server.application;
using amenities_server.model;
using amenities_server.persistence;

namespace amenities_server.validators
{
    public static class AmenityValidator
    {
        private static IAmenityPersistence amenityPersistence;

        /*
        Validates the parameters of an amenity, checking for null or empty values for its Name and Description.
        Args:
            amenity: An instance of the Amenity class to validate.
        Returns:
            bool: Returns true if the amenity has a valid name and description, false otherwise.
        */
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

        /*
        Validates a new amenity by checking its parameters and ensuring that the amenity is either new (id < 0)
        or does not already exist in the persistence layer.
        Args:
            amenity: An instance of the Amenity class to validate.
        Returns:
            bool: Returns true if the amenity parameters are valid and the amenity does not already exist, false otherwise.
        */
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
