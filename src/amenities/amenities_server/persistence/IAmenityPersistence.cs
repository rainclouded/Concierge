public interface IAmenityPersistence
{
    IEnumerable<Amenity> GetAmenities();
    Amenity GetAmenityByID(int id);
    void AddAmenity(Amenity amenity);
    void UpdateAmenity(int id, Amenity amenity);

    void DeleteAmenity(int id);
}
