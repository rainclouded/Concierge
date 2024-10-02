public class StubAmenityPersistence : IAmenityPersistence
{
    private List<Amenity> _amenities;
    private static int _nextId = 1;
    public StubAmenityPersistence()
    {
        _amenities = new List<Amenity>();

        AddAmenity(new Amenity("Pool", "Outdoor pool", new TimeSpan(9, 0, 0), new TimeSpan(21, 0, 0)));
        AddAmenity(new Amenity("Gym", "24/7 access gym", new TimeSpan(0, 0, 0), new TimeSpan(24, 0, 0)));
        AddAmenity(new Amenity("Breakfast", "Free breakfast", new TimeSpan(6, 0, 0), new TimeSpan(10, 0, 0)));
        AddAmenity(new Amenity("Bar", "Serves alcohol and food", new TimeSpan(17, 0, 0), new TimeSpan(2, 0, 0)));
    }

    public IEnumerable<Amenity> GetAmenities(){
        return _amenities;
    }

    public Amenity GetAmenityByID(int id)
    {
        return _amenities.FirstOrDefault(a => a.Id == id);
    }
    
    public Amenity AddAmenity(Amenity amenity)
    {
        if (GetAmenityByID(amenity.Id) != null)
        {
            throw new InvalidOperationException("Amenity with the same ID already exists.");
        }

        amenity.Id = _nextId++;
        _amenities.Add(amenity);

        return amenity;
    }

    public Amenity UpdateAmenity(int id, Amenity amenity)
    {
        var existingAmenity = GetAmenityByID(id);
        
        if (existingAmenity == null)
        {
            throw new KeyNotFoundException("Amenity not found.");
        }

        existingAmenity.UpdateAmenity(amenity);

        return existingAmenity;
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
