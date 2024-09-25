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

            if (Environment.GetEnvironmentVariable("forProduction").Equals("test"))
            {
                _amenityPersistence = new StubAmenityPersistence();
            }
            else
            {
                _amenityPersistence = null;
            }

            _controller = new AmenitiesController(_amenityPersistence);
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

            Assert.IsInstanceOf<Amenity>(_controller.GetAmenityByID(amenity.AmenityID));
        }

        [Test]
        public void GetAmenityByID_InvalidID_ReturnsNotFound()
        {
            //assuming amenities exist within database
            int invalidID = -1;

            Assert.IsInstanceOf<NotFoundObjectResult>(_controller.GetAmenityByID(invalidID));
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

            Assert.IsInstanceOf<Amenity>(_controller.GetAmenityByID(_testValidAmenity.AmenityID));
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

            Assert.IsInstanceOf<NotFoundObjectResult>(_controller.GetAmenityByID(_testInvalidAmenity.AmenityID));
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

            Assert.Equals(_controller.GetAmenityByID(_testValidAmenity.AmenityID), _testUpdatedValidAmenity);
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

            Assert.IsInstanceOf<NotFoundObjectResult>(_controller.GetAmenityByID(_testInvalidAmenity.AmenityID));
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
    }
}