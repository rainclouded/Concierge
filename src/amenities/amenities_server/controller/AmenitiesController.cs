using amenities_server.application;
using amenities_server.model;
using amenities_server.persistence;
using amenities_server.validators;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.DependencyInjection;

namespace amenities_server.Controllers;

[ApiController]
[Route("amenities")]
public class AmenitiesController : ControllerBase
{
    private IAmenityPersistence _amenityPersistence;
    private IPermissionValidator _permissionValidator;
    public AmenitiesController(IPermissionValidator permissionValidator)
    {
        _amenityPersistence = Services.GetAmenityPersistence();
        _permissionValidator = permissionValidator;
    }

    //get: /amenities
    [HttpGet]
    public IActionResult GetAmenities()
    {
        //validate permissions of requester
        if (!Request.Headers.TryGetValue("X-Api-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.VIEW_AMENITES, apiKey!))
        {
            return Unauthorized(new AmenityResponse<int>(ResponseMessages.UNAUTHORIZED, 0));
        }

        var amenities = _amenityPersistence.GetAmenities();

        if (amenities == null)
        {
            return NotFound(new AmenityResponse<string>(ResponseMessages.GET_AMENITIES_FAILED, ""));
        }

        return Ok(new AmenityResponse<IEnumerable<Amenity>>(ResponseMessages.GET_AMENITIES_SUCCESS, amenities));
    }

    //get: /amenities/{id}
    [HttpGet("{id}")]
    public IActionResult GetAmenityByID(int id)
    {
        //validate permissions of requester
        if (!Request.Headers.TryGetValue("X-Api-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.VIEW_AMENITES, apiKey!))
        {
            return Unauthorized(new AmenityResponse<int>(ResponseMessages.UNAUTHORIZED, 0));
        }

        var amenity = _amenityPersistence.GetAmenityByID(id);

        //validate passed id
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
        //validate permissions of requester
        if (!Request.Headers.TryGetValue("X-Api-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.DELETE_AMENITES, apiKey!))
        {
            return Unauthorized(new AmenityResponse<int>(ResponseMessages.UNAUTHORIZED, id));
        }

        //validate if id is valid
        if(_amenityPersistence.GetAmenityByID(id) == null)
        {
            return NotFound(new AmenityResponse<int>(ResponseMessages.GET_AMENITY_FAILED, id));
        }

        _amenityPersistence.DeleteAmenity(id);

        return Ok(new AmenityResponse<string>(ResponseMessages.DELETE_AMENITY_SUCCESS, ""));
    }

    //post: /amenities
    [HttpPost]
    public IActionResult AddAmenity(Amenity newAmenity)
    {
        //validate permissions of requester
        if (!Request.Headers.TryGetValue("X-Api-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.CREATE_AMENITES, apiKey!))
        {
            return Unauthorized(new AmenityResponse<int>(ResponseMessages.UNAUTHORIZED, 0));
        }

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
        //validate permissions of requester
        if (!Request.Headers.TryGetValue("X-Api-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.EDIT_AMENITES, apiKey!))
        {
            return Unauthorized(new AmenityResponse<int>(ResponseMessages.UNAUTHORIZED, 0));
        }

        //validate passed id
        if (_amenityPersistence.GetAmenityByID(id) == null)
        {
            return NotFound(new AmenityResponse<Amenity>(ResponseMessages.GET_AMENITY_FAILED, newAmenity));
        }

        //validate passed amenity
        if (!AmenityValidator.ValidateAmenityParameters(newAmenity))
        {
            return BadRequest(new AmenityResponse<Amenity>(ResponseMessages.INVALID_AMENITY_PASSED, newAmenity));
        }

        newAmenity = _amenityPersistence.UpdateAmenity(id, newAmenity);

        return Ok(new AmenityResponse<Amenity>(ResponseMessages.UPDATE_AMENITY_SUCCESS, newAmenity));
    }
}


