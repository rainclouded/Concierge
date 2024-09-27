using Microsoft.AspNetCore.Mvc;

namespace amenities_server.Controllers;

[ApiController]
[Route("amenities")]
public class AmenitiesController : ControllerBase
{
    private IAmenityPersistence _amenityPersistence;

    public AmenitiesController()
    {
        _amenityPersistence = Services.GetAmenityPersistence();
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

    //delete: /amenities/{id}
    [HttpDelete("{id}")]
    public IActionResult DeleteAmenity(int id)
    {
        //validate session

        //validate passed amenity
        var amenity = _amenityPersistence.GetAmenityByID(id);
        if(amenity == null){
            return BadRequest("Invalid request! An non existent amenity was requested.");
        }

        _amenityPersistence.DeleteAmenity(id);
        return Ok("Amenity deleted successfully.");
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

    //put: /amenities/{id}
    [HttpPut("{id}")]
    public IActionResult UpdateAmenity(int id, Amenity newAmenity)
    {
        //validate session

        if(_amenityPersistence.GetAmenityByID(id) == null)
        {
            return BadRequest("Invalid amenity! A non existent amenity was requested to be updated");
        }

        if (!AmenityValidator.ValidateExistingAmenity(newAmenity))
        {
            return BadRequest("Invalid amenity! An amenity with invalid parameters was passed.");
        }

        _amenityPersistence.UpdateAmenity(id, newAmenity);

        return Ok("Amenity updated successfully");
    }
}


