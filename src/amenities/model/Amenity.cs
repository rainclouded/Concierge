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
}
