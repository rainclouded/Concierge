using Microsoft.AspNetCore.Mvc;

namespace amenities_server.Controllers;

[ApiController]
[Route("amenities")]
public class AmenitiesController : ControllerBase
{
    private IAmenityPersistence _amenityPersistence;

    public AmenitiesController(IAmenityPersistence amenityPersistence = null)
    {
        if (amenityPersistence == null)
        {
            _amenityPersistence = Services.GetAmenityPersistence(Environment.GetEnvironmentVariable("forProduction"));
        }
        else
        {
            _amenityPersistence = amenityPersistence;
        }
    }

    //get: /amenities
    [HttpGet]
    public IActionResult GetAmenities()
    {
        var amenities = _amenityPersistence.GetAmenities();

        if (amenities == null)
        {
            return NotFound();
        }

        return Ok(amenities);
    }

    //get: /amenities/{id}
    [HttpGet("{id}")]
    public ActionResult<Amenity> GetAmenityByID(int id)
    {
        var amenity = _amenityPersistence.GetAmenityByID(id);

        if (amenity == null)
        {
            return NotFound();
        }

        return amenity;
    }

    //post: /amenities
    [HttpPost]
    public IActionResult AddAmenity(Amenity newAmenity)
    {
        //validate session

        //validate passed amenity
        if (!AmenityValidator.ValidateNewAmenity(newAmenity))
        {
            return BadRequest("Invalid amenity! An amenity with invalid parameters was passed.");
        }

        _amenityPersistence.AddAmenity(newAmenity);

        return Ok("Amenity added successfully.");
    }

    //put: /amenities
    [HttpPut]
    public IActionResult UpdateAmenity(Amenity newAmenity)
    {
        //validate session

        if (!AmenityValidator.ValidateExistingAmenity(newAmenity))
        {
            return BadRequest("Invalid amenity! An amenity with invalid parameters was passed.");
        }

        _amenityPersistence.UpdateAmenity(newAmenity);

        return Ok("Amenity updated successfully");
    }
}


