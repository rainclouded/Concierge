using amenities_server;
using amenities_server.application;
using amenities_server.model;
using amenities_server.persistence;
using Microsoft.AspNetCore.Http.HttpResults;
using Npgsql;

namespace amenities_db_integration_test
{
    public class DbInterfaceTests
    {
        private PostgresAmenityPersistence? _persistence;
        private string? _connectionString;

        [OneTimeSetUp]
        public void OneTimeSetUp() 
        {   
            _connectionString = "Host=127.0.0.1; Port=50014; Username=postgres; Password=sa";

            _persistence = new PostgresAmenityPersistence(_connectionString);
            Services.SetAmenityPersistence(_persistence);
        }

        [SetUp]
        public void Setup()
        {
            using var connection = new NpgsqlConnection("Host=127.0.0.1; Port=50014; Username=postgres; Password=sa");
            connection.Open();
            using var command = new NpgsqlCommand("delete from amenities", connection);
            command.ExecuteReader();
        }

        [Test]
        public void GetEmptyAmenities()
        {
            Assert.That(_persistence!.GetAmenities().Count(), Is.EqualTo(0));
        }

        [Test]
        public void GetEmptyAmenityById()
        {
            Assert.That(_persistence!.GetAmenityByID(0), Is.Null);
        }

        [Test]
        public void AddToEmptyAmenityById()
        {
            var created = _persistence!.AddAmenity(NewAmenity());
            var found = _persistence!.GetAmenityByID(created.Id);
            Assert.That(AreAmenitiesEqual(created, found), Is.True);
        }

        [Test]
        public void DeleteEmptyAmenityById()
        {
            Assert.DoesNotThrow(() => _persistence!.DeleteAmenity(1));
        }

        [Test]
        public void UpdateEmptyAmenityById()
        {
            Assert.That(_persistence!.UpdateAmenity(0, NewAmenity()), Is.Null);
        }

        [Test]
        public void GetOneAmenities()
        {
            var amenity = AddAmenity();
            var foundAmenities = _persistence!.GetAmenities();
            Assert.That(foundAmenities.Count(), Is.EqualTo(1));
            Assert.That(AreAmenitiesEqual(amenity, foundAmenities.First()), Is.True);
        }

        [Test]
        public void GetOneAmenityById()
        {
            var amenity = AddAmenity();
            var foundAmenity = _persistence!.GetAmenityByID(amenity.Id);
            Assert.That(AreAmenitiesEqual(amenity, foundAmenity), Is.True);
        }

        [Test]
        public void GetOneAmenityByIdNotFound()
        {
            var amenity = AddAmenity();
            var foundAmenity = _persistence!.GetAmenityByID(0);
            Assert.That(foundAmenity, Is.Null);
        }

        [Test]
        public void AddToOneAmenityById()
        {
            var amenity = AddAmenity();
            var created = _persistence!.AddAmenity(NewAmenity());
            var found = _persistence!.GetAmenityByID(created.Id);
            Assert.That(AreAmenitiesEqual(created, found), Is.True);
        }

        [Test]
        public void DeleteOneAmenityById()
        {
            var amenity = AddAmenity();
            Assert.DoesNotThrow(() => _persistence!.DeleteAmenity(amenity.Id));
            Assert.That(_persistence!.GetAmenities().Count(), Is.EqualTo(0));
        }

        [Test]
        public void DeleteOneAmenityByIdNotFound()
        {
            var amenity = AddAmenity();
            Assert.DoesNotThrow(() => _persistence!.DeleteAmenity(-1));
            Assert.That(_persistence!.GetAmenities().Count(), Is.EqualTo(1));
        }

        [Test]
        public void UpdateOneAmenityById()
        {
            var amenity = AddAmenity();
            amenity.Name = "Cat";
            amenity.Description = "Car";
            amenity.StartTime = new TimeSpan(3, 1, 1);
            amenity.EndTime = new TimeSpan(4, 1, 1);
            var newAmenity = _persistence!.UpdateAmenity(amenity.Id, amenity);
            Assert.That(AreAmenitiesEqual(newAmenity, amenity), Is.True);
            Assert.That(AreAmenitiesEqual(newAmenity, _persistence.GetAmenityByID(amenity.Id)), Is.True);
        }

        [Test]
        public void UpdateOneAmenityByIdNotFound()
        {
            var amenity = AddAmenity();
            Assert.That(_persistence!.UpdateAmenity(-1, NewAmenity()), Is.Null);
        }


        private Amenity AddAmenity(string name = "test name", string description = "test description", TimeSpan? start = null, TimeSpan? end = null)
        {
            var amenity = new Amenity(0, name, description, start ?? new TimeSpan(1, 0, 0), end ?? new TimeSpan(2, 0, 0));
            using var connection = new NpgsqlConnection(_connectionString);
            connection.Open();
            using var command = new NpgsqlCommand("INSERT INTO amenities (name, description, start_time, end_time) values (@name, @description, @startTime, @endTime) RETURNING id;", connection);
            command.Parameters.AddWithValue("name", amenity.Name);
            command.Parameters.AddWithValue("description", amenity.Description);
            command.Parameters.AddWithValue("startTime", amenity.StartTime);
            command.Parameters.AddWithValue("endTime", amenity.EndTime);
            amenity.Id = (int)(command.ExecuteScalar() ?? -1);
            return amenity;
        }

        private static Amenity NewAmenity(string name = "test name", string description = "test description", TimeSpan? start = null, TimeSpan? end = null)
        {
            return new Amenity(0, name, description, start ?? new TimeSpan(1, 0, 0), end ?? new TimeSpan(2, 0, 0));
        }

        private static bool AreAmenitiesEqual(Amenity a, Amenity b)
        {
            return a.Id == b.Id && a.Name == b.Name && a.Description == b.Description && a.StartTime.Equals(b.StartTime) && a.EndTime.Equals(b.EndTime);
        }
    }
}