using amenities_server.Controllers;
using Microsoft.AspNetCore.Mvc;

namespace amenities_test
{
    public class AmenitiesControllerTests
    { 
        private AmenitiesController _controller;
        private IAmenityPersistence _amenityPersistence;
        private Amenity _testValidAmenity;
        private Amenity _testUpdatedValidAmenity;
        private Amenity _testInvalidAmenity;


        [SetUp]
        public void Setup()
        {
            _testValidAmenity = new Amenity(999, "testValidAmenity", "testValidDesc", new TimeSpan(0, 0, 0), new TimeSpan(12, 0, 0));
            _testUpdatedValidAmenity = new Amenity(999, "testInvalidAmenity", "testValidDesc", new TimeSpan(12, 0, 0), new TimeSpan(24, 0, 0));
            _testInvalidAmenity = new Amenity(999, "testInvalidAmenity", "testValidDesc", new TimeSpan(13, 0, 0), new TimeSpan(12, 0, 0));

            Services.clearPersistences();
            _amenityPersistence = Services.GetAmenityPersistence();
  
            _controller = new AmenitiesController();
        }

        [Test]
        public void GetAmenities_ShouldNotReturnNull()
        {
            var result = _controller.GetAmenities();

            Assert.NotNull(result);
            Assert.IsInstanceOf<OkObjectResult>(result);
        }

        [Test]
        public void GetAmenityByID_ValidID_NotReturnNull()
        {
            _amenityPersistence.AddAmenity(_testValidAmenity);
            var amenity = _amenityPersistence.GetAmenityByID(_testValidAmenity.AmenityID);

            Assert.IsInstanceOf<Amenity>(_controller.GetAmenityByID(amenity.AmenityID).Value);
        }

        [Test]
        public void GetAmenityByID_InvalidID_ReturnsNotFound()
        {
            //assuming amenities exist within database
            int invalidID = -1;

            Assert.IsInstanceOf<NotFoundResult>(_controller.GetAmenityByID(invalidID).Result);
        }

        [Test]
        public void AddAmenity_ValidAmenity_IsSuccessful()
        {
            Assert.IsInstanceOf<OkObjectResult>(_controller.AddAmenity(_testValidAmenity));
        }

        [Test]
        public void AddAmenity_ValidAmenity_AbleToBeFetched()
        {
            _controller.AddAmenity(_testValidAmenity);

            Assert.IsInstanceOf<Amenity>(_controller.GetAmenityByID(_testValidAmenity.AmenityID).Value);
        }
        [Test]
        public void AddAmenity_InvalidAmenity_Fails()
        {
            Assert.IsInstanceOf<BadRequestObjectResult>(_controller.AddAmenity(_testInvalidAmenity));
        }

        [Test]
        public void AddAmenity_InvalidAmenity_NotFetchable()
        {
            _controller.AddAmenity(_testInvalidAmenity);

            Assert.IsInstanceOf<NotFoundResult>(_controller.GetAmenityByID(_testInvalidAmenity.AmenityID).Result);
        }

        [Test]
        public void AddAmenity_DuplicateAmenity_Fails()
        {
            _amenityPersistence.AddAmenity(_testValidAmenity);

            Assert.IsInstanceOf<BadRequestObjectResult>(_controller.AddAmenity(_testValidAmenity));
        }

        [Test]
        public void AddAmenity_NullAmenity_Fails()
        {
            Assert.IsInstanceOf<BadRequestObjectResult>(_controller.AddAmenity(null));
        }
        [Test]
        public void UpdateAmenity_ValidAmenity_IsSuccessful()
        {
            _amenityPersistence.AddAmenity(_testValidAmenity);
            Assert.IsInstanceOf<OkObjectResult>(_controller.UpdateAmenity(_testValidAmenity));

        }

        [Test]
        public void UpdateAmenity_ValidAmenity_UpdatedAmenityFetched()
        {
            _amenityPersistence.AddAmenity(_testValidAmenity);

            _controller.UpdateAmenity(_testUpdatedValidAmenity);

            Assert.IsTrue(_controller.GetAmenityByID(_testValidAmenity.AmenityID).Value.Equals(_testUpdatedValidAmenity));
        }
        [Test]
        public void UpdateAmenity_InvalidAmenity_Fails()
        {
            Assert.IsInstanceOf<BadRequestObjectResult>(_controller.UpdateAmenity(_testInvalidAmenity));
        }

        [Test]
        public void UpdateAmenity_InvalidAmenity_NotFetchable()
        {
            _controller.UpdateAmenity(_testInvalidAmenity);

            Assert.IsInstanceOf<NotFoundResult>(_controller.GetAmenityByID(_testInvalidAmenity.AmenityID).Result);
        }

        [Test]
        public void UpdateAmenity_NonExistingAmenity_Fails()
        {
            Assert.IsInstanceOf<BadRequestObjectResult>(_controller.UpdateAmenity(_testValidAmenity));
        }

        [Test]
        public void UpdateAmenity_NullAmenity_Fails()
        {
            Assert.IsInstanceOf<BadRequestObjectResult>(_controller.UpdateAmenity(null));
        }

        [Test]
        public void DeleteAmenity_ValidAmenity_IsSuccessful()
        {
            _amenityPersistence.AddAmenity(_testValidAmenity);

            Assert.IsInstanceOf<OkObjectResult>(_controller.DeleteAmenity(_testValidAmenity.AmenityID));

        }

        [Test]
        public void DeleteAmenity_ValidAmenity_UpdatedAmenityNotFetched()
        {
            _amenityPersistence.AddAmenity(_testValidAmenity);

            _controller.DeleteAmenity(_testValidAmenity.AmenityID);

            Assert.IsInstanceOf<NotFoundResult>(_controller.GetAmenityByID(_testValidAmenity.AmenityID).Result);
        }
        [Test]
        public void DeleteAmenity_InvalidAmenity_Fails()
        {
            Assert.IsInstanceOf<BadRequestObjectResult>(_controller.DeleteAmenity(_testInvalidAmenity.AmenityID));
        }

        [Test]
        public void DeleteAmenity_InvalidAmenity_NotFetchable()
        {
            _controller.DeleteAmenity(_testInvalidAmenity.AmenityID);

            Assert.IsInstanceOf<NotFoundResult>(_controller.GetAmenityByID(_testInvalidAmenity.AmenityID).Result);
        }
    }
}