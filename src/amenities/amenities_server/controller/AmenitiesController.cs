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

        return Ok(new AmenityResponse<IEnumerable<Amenity>>("Amenities retrieved successfully.", amenities));
    }

    //get: /amenities/{id}
    [HttpGet("{id}")]
    public IActionResult GetAmenityByID(int id)
    {
        var amenity = _amenityPersistence.GetAmenityByID(id);

        if (amenity == null)
        {
            return NotFound(new AmenityResponse<int>("Amenity with specified id not found.", id));
        }

        return Ok(new AmenityResponse<Amenity>("Amenity retrieved successfully.", amenity));
    }

    //delete: /amenities/{id}
    [HttpDelete("{id}")]
    public IActionResult DeleteAmenity(int id)
    {
        //TODO:  validate session call
        //if(!Services.GetPermissionValidator().ValidatePermissions(permission,sessionKey))

        //validate passed amenity
        var amenity = _amenityPersistence.GetAmenityByID(id);
        if(amenity == null){
            return BadRequest(new AmenityResponse<int>("Bad Request. Amenity with specified id not found.", id));
        }

        _amenityPersistence.DeleteAmenity(id);

        return Ok(new AmenityResponse<string>("Amenity deleted successfully.", null));
    }
    //post: /amenities
    [HttpPost]
    public IActionResult AddAmenity(Amenity newAmenity)
    {
        //TODO: validate session call
        //if(!Services.GetPermissionValidator().ValidatePermissions(permission,sessionKey))

        //validate passed amenity
        if (!AmenityValidator.ValidateNewAmenity(newAmenity))
        {
            return BadRequest(new AmenityResponse<Amenity>("Bad Request. Amenity with invalid parameters was passed.", newAmenity));
        }

        _amenityPersistence.AddAmenity(newAmenity);

        //create uri that points to the newly added amenity
        var uri = $"{Request.Scheme}://{Request.Host}/amenities/{newAmenity.Id}";
        return Created(uri, new AmenityResponse<Amenity>("Amenity created successfully.", newAmenity));
    }

    //put: /amenities/{id}
    [HttpPut("{id}")]
    public IActionResult UpdateAmenity(int id, Amenity newAmenity)
    {
        //TODO:  validate session call
        //if(!Services.GetPermissionValidator().ValidatePermissions(permission,sessionKey))

        if(_amenityPersistence.GetAmenityByID(id) == null)
        {
            return BadRequest(new AmenityResponse<Amenity>("Bad Request. Non existing amenity was requested to be updated.", newAmenity));
        }

        if (!AmenityValidator.ValidateAmenityParameters(newAmenity))
        {
            return BadRequest(new AmenityResponse<Amenity>("Bad Request. Amenity with invalid parameters was passed.", newAmenity));
        }

        _amenityPersistence.UpdateAmenity(id, newAmenity);

        return Ok(new AmenityResponse<Amenity>("Amenity was updated successfully.", newAmenity));
    }
}


