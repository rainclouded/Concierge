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
        public Amenity(int id, string name, string description, TimeSpan startTime, TimeSpan endTime)
        {
            Id = id;
            Name = name;
            Description = description;
            StartTime = startTime;
            EndTime = endTime;
        }

        public Amenity(string name, string description, TimeSpan startTime, TimeSpan endTime)
        {
            Name = name;
            Description = description;
            StartTime = startTime;
            EndTime = endTime;
        }

        public void UpdateAmenity(Amenity updatedAmenity)
        {
            Name = updatedAmenity.Name;
            Description = updatedAmenity.Description;
            StartTime = updatedAmenity.StartTime;
            EndTime = updatedAmenity.EndTime;
        }

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

        public override int GetHashCode()
        {
            return Id;
        }
    }
}