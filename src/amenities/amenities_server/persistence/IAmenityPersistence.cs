using amenities_server.model;

namespace amenities_server.persistence
{
    public interface IAmenityPersistence
    {
        IEnumerable<Amenity> GetAmenities();
        Amenity GetAmenityByID(int id);
        Amenity AddAmenity(Amenity amenity);
        Amenity UpdateAmenity(int id, Amenity amenity);

        void DeleteAmenity(int id);
    }
}