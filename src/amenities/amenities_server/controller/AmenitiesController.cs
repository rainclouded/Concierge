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

    /*
    Initializes the AmenitiesController with the provided permission validator.
    Args:
        permissionValidator: A permission validator to check the requester's permissions for different actions
    Returns:
        None
    */
    public AmenitiesController(IPermissionValidator permissionValidator)
    {
        _amenityPersistence = Services.GetAmenityPersistence();
        _permissionValidator = permissionValidator;
    }

    //get: /amenities
    /*
    Retrieves a list of all amenities. Validates the requester's permissions before fetching the data.
    Args:
        None
    Returns:
        IActionResult: Returns Ok with a list of amenities if the request is successful, 
        or Unauthorized/NotFound if permissions or data issues occur.
    */
    [HttpGet]
    public IActionResult GetAmenities()
    {
        //validate permissions of requester
        if (!Request.Headers.TryGetValue("X-API-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.VIEW_AMENITES, apiKey!))
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
    /*
    Retrieves a specific amenity by its ID. Verifies the requester's permissions and checks if the requested amenity exists.
    Args:
        id: The unique identifier of the amenity to retrieve
    Returns:
        IActionResult: Returns Ok with the requested amenity if found, or 
        Unauthorized/NotFound if there are permission issues or the amenity does not exist.
    */
    [HttpGet("{id}")]
    public IActionResult GetAmenityByID(int id)
    {
        //validate permissions of requester
        if (!Request.Headers.TryGetValue("X-API-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.VIEW_AMENITES, apiKey!))
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
    /*
    Deletes a specific amenity by its ID after validating the requester's
     permissions. Verifies if the amenity exists before attempting to delete it.
    Args:
        id: The unique identifier of the amenity to delete
    Returns:
        IActionResult: Returns Ok if the amenity is successfully 
        deleted, or Unauthorized/NotFound if there are permission issues or the amenity does not exist.
    */
    [HttpDelete("{id}")]
    public IActionResult DeleteAmenity(int id)
    {
        //validate permissions of requester
        if (!Request.Headers.TryGetValue("X-API-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.DELETE_AMENITES, apiKey!))
        {
            return Unauthorized(new AmenityResponse<int>(ResponseMessages.UNAUTHORIZED, id));
        }

        //validate if id is valid
        if (_amenityPersistence.GetAmenityByID(id) == null)
        {
            return NotFound(new AmenityResponse<int>(ResponseMessages.GET_AMENITY_FAILED, id));
        }

        _amenityPersistence.DeleteAmenity(id);

        return Ok(new AmenityResponse<string>(ResponseMessages.DELETE_AMENITY_SUCCESS, ""));
    }

    //post: /amenities
    /*
    Adds a new amenity to the system after validating the requester's permissions and the input data.
    Args:
        newAmenity: The new amenity object to be added to the system
    Returns:
        IActionResult: Returns CreatedAtAction with the newly created 
        amenity if successful, or Unauthorized/BadRequest if there are permission issues or invalid data.
    */
    [HttpPost]
    public IActionResult AddAmenity(Amenity newAmenity)
    {
        //validate permissions of requester
        if (!Request.Headers.TryGetValue("X-API-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.EDIT_AMENITES, apiKey!))
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
    /*
    Updates an existing amenity by its ID. Validates the requester's 
    permissions, checks if the amenity exists, and validates the input data before updating.
    Args:
        id: The ID of the amenity to update
        newAmenity: The updated amenity data to replace the existing one
    Returns:
        IActionResult: Returns Ok with the updated amenity if successful, 
        or Unauthorized/NotFound/BadRequest if there are permission issues, missing data, or validation errors.
    */
    [HttpPut("{id}")]
    public IActionResult UpdateAmenity(int id, Amenity newAmenity)
    {
        //validate permissions of requester
        if (!Request.Headers.TryGetValue("X-API-Key", out var apiKey) || !_permissionValidator.ValidatePermissions(PermissionNames.EDIT_AMENITES, apiKey!))
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
