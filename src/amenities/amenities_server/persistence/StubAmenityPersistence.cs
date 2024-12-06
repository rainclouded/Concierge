using amenities_server.model;

namespace amenities_server.persistence
{
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

        /*
        Retrieves all amenities from the in-memory list.
        Returns:
            IEnumerable<Amenity>: A collection of all the amenities stored in memory.
        */
        public IEnumerable<Amenity> GetAmenities()
        {
            return _amenities;
        }

        /*
        Retrieves a specific Amenity by its ID from the in-memory list.
        Args:
            id: The ID of the Amenity to retrieve.
        Returns:
            Amenity: The Amenity object that matches the provided ID, or null if not found.
        */
        public Amenity GetAmenityByID(int id)
        {
            return _amenities.FirstOrDefault(a => a.Id == id);
        }

        /*
        Adds a new Amenity to the in-memory list with a unique ID.
        Throws an exception if an Amenity with the same ID already exists.
        Args:
            amenity: The Amenity object that is to be added to the list.
        Returns:
            Amenity: The added Amenity with its assigned ID.
        */
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

        /*
        Updates an existing Amenity in the in-memory list.
        Throws an exception if the Amenity to update does not exist.
        Args:
            id: The ID of the Amenity to update.
            amenity: The new Amenity object with updated values.
        Returns:
            Amenity: The updated Amenity object.
        */
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

        /*
        Deletes an Amenity from the in-memory list based on its ID.
        Throws an exception if the Amenity to delete does not exist.
        Args:
            id: The ID of the Amenity to delete.
        */
        public void DeleteAmenity(int id)
        {
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
}
