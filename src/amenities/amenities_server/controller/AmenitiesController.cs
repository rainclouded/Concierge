using Microsoft.AspNetCore.Mvc;

namespace amenities_server.Controllers;

[ApiController]
[Route("amenities")]
public class AmenitiesController : ControllerBase
{
    private IAmenityPersistence _amenityPersistence;
    //private IPermissionValidator _permissionValidator
    public AmenitiesController()
    {
        _amenityPersistence = Services.GetAmenityPersistence();
        //_permissionValidator = Services.GetPermissionValidator();
    }

    //get: /amenities
    [HttpGet]
    public IActionResult GetAmenities()
    {
        var amenities = _amenityPersistence.GetAmenities();

        if (amenities == null)
        {
            return NotFound(new AmenityResponse<string>(ResponseMessages.GET_AMENITIES_FAILED, null));
        }

        return Ok(new AmenityResponse<IEnumerable<Amenity>>(ResponseMessages.GET_AMENITIES_SUCCESS, amenities));
    }

    //get: /amenities/{id}
    [HttpGet("{id}")]
    public IActionResult GetAmenityByID(int id)
    {
        var amenity = _amenityPersistence.GetAmenityByID(id);

        if (_amenityPersistence.GetAmenityByID(id) == null)
        {
            return NotFound(new AmenityResponse<int>(ResponseMessages.GET_AMENITY_FAILED, id));
        }

        return Ok(new AmenityResponse<Amenity>(ResponseMessages.GET_AMENITY_SUCCESS, amenity));
    }

    //delete: /amenities/{id}
    [HttpDelete("{id}")]
    public IActionResult DeleteAmenity(int id)
    {
        //TODO: validate session call
        //if(_permissionValidator.ValidatePermissions(permission,sessionKey))

        //validate if id is valid
        if(_amenityPersistence.GetAmenityByID(id) == null)
        {
            return NotFound(new AmenityResponse<int>(ResponseMessages.GET_AMENITY_FAILED, id));
        }

        _amenityPersistence.DeleteAmenity(id);

        return Ok(new AmenityResponse<string>(ResponseMessages.DELETE_AMENITY_SUCCESS, null));
    }
    //post: /amenities
    [HttpPost]
    public IActionResult AddAmenity(Amenity newAmenity)
    {
        //TODO: validate session call
        //if(_permissionValidator.ValidatePermissions(permission,sessionKey))
        // 
        //validate passed amenity
        if (!AmenityValidator.ValidateNewAmenity(newAmenity))
        {
            return BadRequest(new AmenityResponse<Amenity>(ResponseMessages.INVALID_AMENITY_PASSED, newAmenity));
        }

        _amenityPersistence.AddAmenity(newAmenity);

        //return a 201 with location to newly created amenity
        return CreatedAtAction(
            nameof(GetAmenityByID),  
            new { id = newAmenity.Id },  
            new AmenityResponse<Amenity>(ResponseMessages.CREATE_AMENITY_SUCCESS, newAmenity)
        );
    }

    //put: /amenities/{id}
    [HttpPut("{id}")]
    public IActionResult UpdateAmenity(int id, Amenity newAmenity)
    {
        //TODO:  validate session call
        //if(!Services.GetPermissionValidator().ValidatePermissions(permission,sessionKey))

        if(_amenityPersistence.GetAmenityByID(id) == null)
        {
            return NotFound(new AmenityResponse<Amenity>(ResponseMessages.GET_AMENITY_FAILED, newAmenity));
        }

        if (!AmenityValidator.ValidateAmenityParameters(newAmenity))
        {
            return BadRequest(new AmenityResponse<Amenity>(ResponseMessages.INVALID_AMENITY_PASSED, newAmenity));
        }

        newAmenity = _amenityPersistence.UpdateAmenity(id, newAmenity);

        return Ok(new AmenityResponse<Amenity>(ResponseMessages.UPDATE_AMENITY_SUCCESS, newAmenity));
    }
}


