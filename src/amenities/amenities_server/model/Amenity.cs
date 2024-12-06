namespace amenities_server.model
{
    public class Amenity
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public string Description { get; set; }
        public TimeSpan StartTime { get; set; }
        public TimeSpan EndTime { get; set; }


        public Amenity() { }

        /*
        Constructor that initializes an Amenity object with specific values.
        Args:
            id: The unique identifier of the amenity
            name: The name of the amenity
            description: A description of the amenity
            startTime: The start time when the amenity is available
            endTime: The end time when the amenity is available
        Returns:
            None
        */
        public Amenity(int id, string name, string description, TimeSpan startTime, TimeSpan endTime)
        {
            Id = id;
            Name = name;
            Description = description;
            StartTime = startTime;
            EndTime = endTime;
        }

        /*
        Constructor that initializes an Amenity object without the id.
        Args:
            name: The name of the amenity
            description: A description of the amenity
            startTime: The start time when the amenity is available
            endTime: The end time when the amenity is available
        Returns:
            None
        */
        public Amenity(string name, string description, TimeSpan startTime, TimeSpan endTime)
        {
            Name = name;
            Description = description;
            StartTime = startTime;
            EndTime = endTime;
        }

        /*
        Updates the current amenity with the values from another Amenity object.
        Args:
            updatedAmenity: The Amenity object containing the new values for updating the current Amenity
        Returns:
            None
        */
        public void UpdateAmenity(Amenity updatedAmenity)
        {
            Name = updatedAmenity.Name;
            Description = updatedAmenity.Description;
            StartTime = updatedAmenity.StartTime;
            EndTime = updatedAmenity.EndTime;
        }

        /*
        Compares two Amenity objects for equality based on their Id, Name, Description, StartTime, and EndTime.
        Args:
            obj: The object to compare against
        Returns:
            bool: True if the two Amenity objects are equal, false otherwise
        */
        public override bool Equals(object? obj)
        {
            if (obj is Amenity otherAmenity)
            {
                return Id == otherAmenity.Id && Name.Equals(otherAmenity.Name) &&
                       Description.Equals(otherAmenity.Description) && StartTime.Equals(otherAmenity.StartTime) &&
                       EndTime.Equals(otherAmenity.EndTime);
            }

            return false;
        }

        /*
        Returns a hash code for the Amenity object
        Args:
            None
        Returns:
            int: The hash code for the Amenity object
        */
        public override int GetHashCode()
        {
            return Id;
        }
    }
}
