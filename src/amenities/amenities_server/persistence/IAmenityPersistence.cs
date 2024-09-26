public interface IAmenityPersistence
{
    IEnumerable<Amenity> GetAmenities();
    Amenity GetAmenityByID(int id);
    void AddAmenity(Amenity amenity);
    void UpdateAmenity(Amenity amenity);

    void DeleteAmenity(int id);
}
