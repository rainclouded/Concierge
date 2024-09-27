public class StubAmenityPersistence : IAmenityPersistence
{
    private List<Amenity> _amenities;

    public StubAmenityPersistence()
    {
        _amenities = new List<Amenity>
        {
            new Amenity(1, "Pool", "Outdoor pool", new TimeSpan(9, 0, 0), new TimeSpan(21, 0, 0)),
            new Amenity(2, "Gym", "24/7 access gym", new TimeSpan(0, 0, 0), new TimeSpan(24, 0, 0)),
            new Amenity(3, "Breakfast", "Free breakfast", new TimeSpan(6, 0, 0), new TimeSpan(10, 0, 0)),
            new Amenity(4, "Bar", "Serves alcohol and food", new TimeSpan(17, 0, 0), new TimeSpan(2, 0, 0))
        };
    }

    public IEnumerable<Amenity> GetAmenities(){
        return _amenities;
    }

    public Amenity GetAmenityByID(int id)
    {
        return _amenities.FirstOrDefault(a => a.Id == id);
    }
    

    public void AddAmenity(Amenity amenity)
    {
        if (_amenities.Any(a => a.Id == amenity.Id))
        {
            throw new InvalidOperationException("Amenity with the same ID already exists.");
        }

        _amenities.Add(amenity);
    }

    public void UpdateAmenity(int id, Amenity amenity)
    {
        var existingAmenity = GetAmenityByID(id);
        
        if (existingAmenity == null)
        {
            throw new KeyNotFoundException("Amenity not found.");
        }

        existingAmenity.updateAmenity(amenity);
    }

    public void DeleteAmenity(int id){
        var amenityToDelete = GetAmenityByID(id);
        
        if (amenityToDelete != null)
        {
            _amenities.Remove(amenityToDelete);
        }
        else
        {
            throw new KeyNotFoundException("Amenity not found.");
        }
    }
}
