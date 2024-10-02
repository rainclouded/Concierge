public class Amenity
{
    public int Id { get; set; }
    public string Name { get; set; }
    public string Description { get; set; }
    public TimeSpan StartTime { get; set; }
    public TimeSpan EndTime { get; set; }

    public Amenity(string name, string description, TimeSpan startTime, TimeSpan endTime)
    {
        Name = name;
        Description = description;
        StartTime = startTime;
        EndTime = endTime;
    }

    public void UpdateAmenity(Amenity updatedAmenity){
        Name = updatedAmenity.Name;
        Description = updatedAmenity.Description;
        StartTime = updatedAmenity.StartTime;
        EndTime = updatedAmenity.EndTime;
    }
    
    public override bool Equals(object? obj)
    {
        if(obj is Amenity otherAmenity)
        {
            return this.Id == otherAmenity.Id && this.Name.Equals(otherAmenity.Name) &&
                   this.Description.Equals(otherAmenity.Description) && this.StartTime.Equals(otherAmenity.StartTime) &&
                   this.EndTime.Equals(otherAmenity.EndTime);
        }

        return false;
    }

    public override int GetHashCode()
    {
        return Id;
    }
}
