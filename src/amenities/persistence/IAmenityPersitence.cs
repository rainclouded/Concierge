public interface IAmenityPersistence
{
    IEnumerable<Amenity> GetAmenities();
    IAmenityPersistence GetAmenityByID(int id);
    void AddAmenity(Amenity amenity);
    void UpdateAmenity(Amenity amenity);
}
