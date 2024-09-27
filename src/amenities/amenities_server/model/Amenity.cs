public class Amenity
{
    public int AmenityID { get; set; }
    public string AmenityName { get; set; }
    public string AmenityDescription { get; set; }
    public TimeSpan StartTime { get; set; }
    public TimeSpan EndTime { get; set; }

    public Amenity(int amenityID, string amenityName, string amenityDescription, TimeSpan startTime, TimeSpan endTime)
    {
        AmenityID = amenityID;
        AmenityName = amenityName;
        AmenityDescription = amenityDescription;
        StartTime = startTime;
        EndTime = endTime;
    }

    public override bool Equals(object? obj)
    {
        if(obj is Amenity otherAmenity)
        {
            return this.AmenityID == otherAmenity.AmenityID && this.AmenityName.Equals(otherAmenity.AmenityName) &&
                   this.AmenityDescription.Equals(otherAmenity.AmenityDescription) && this.StartTime.Equals(otherAmenity.StartTime) &&
                   this.EndTime.Equals(otherAmenity.EndTime);
        }

        return false;
    }

    public override int GetHashCode()
    {
        return AmenityID;
    }
}
