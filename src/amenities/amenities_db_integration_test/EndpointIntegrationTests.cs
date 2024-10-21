using amenities_server.model;
using Microsoft.AspNetCore.Mvc.Testing;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;
using Npgsql;
using System.Text;

namespace amenities_db_integration_test
{
    [TestFixture]
    internal class EndpointIntegrationTests
    {
        private string _connectionString;
        private WebApplicationFactory<Program> _appFactory;
        private HttpClient _client;


        [OneTimeSetUp]
        public void OneTimeSetUp()
        {
            _connectionString = "Host=127.0.0.1; Port=50014; Username=postgres; Password=sa";
        }

        [SetUp]
        public void Setup()
        {
            using var connection = new NpgsqlConnection(_connectionString);
            connection.Open();
            using var command = new NpgsqlCommand("delete from amenities", connection);
            command.ExecuteReader();

            _appFactory = new WebAppFactoryForTests(_connectionString); //Automatically sets the db to test postgres
            _client = _appFactory.CreateClient();
        }

        [Test]
        public async Task GetAmenities_Empty()
        {
            var response = await _client.GetAsync("/amenities");
            var content = JsonConvert.DeserializeObject<dynamic>(await response.Content.ReadAsStringAsync());
            Assert.IsNotNull(content);
            Assert.IsNotNull(content.data);
            Assert.IsInstanceOf<JArray> (content.data);
            Assert.AreEqual(0, content.data.Count);
        }

        [Test]
        public async Task GetAmenities_One()
        {
            var amenity = AddAmenity(NewAmenity());
            var response = await _client.GetAsync("/amenities");
            var content = JsonConvert.DeserializeObject<dynamic>(await response.Content.ReadAsStringAsync());
            Assert.IsNotNull(content);
            Assert.IsNotNull(content.data);
            Assert.IsInstanceOf<JArray>(content.data);
            Assert.AreEqual(1, content.data.Count);
            Assert.That(IsAmenityObjEqualToAmenity(content.data[0], amenity));
        }

        [Test]
        public async Task GetAmenities_Two()
        {
            var amenity1 = AddAmenity(NewAmenity());
            var amenity2 = AddAmenity(NewAmenity(name: "name", description: "desc"));
            var response = await _client.GetAsync("/amenities");
            var content = JsonConvert.DeserializeObject<dynamic>(await response.Content.ReadAsStringAsync());
            Assert.IsNotNull(content);
            Assert.IsNotNull(content.data);
            Assert.IsInstanceOf<JArray>(content.data);
            Assert.AreEqual(2, content.data.Count);
            Assert.That(IsAmenityObjEqualToAmenity(content.data[0], amenity1));
            Assert.That(IsAmenityObjEqualToAmenity(content.data[1], amenity2));
        }

        [Test]
        public async Task GetAmenityById_Empty()
        {
            var response = await _client.GetAsync("/amenities/123");
            Assert.That(response.StatusCode, Is.EqualTo(System.Net.HttpStatusCode.NotFound));
        }

        [Test]
        public async Task GetAmenityById_One()
        {
            var amenity = AddAmenity(NewAmenity());
            var response = await _client.GetAsync($"/amenities/{amenity.Id}");
            var content = JsonConvert.DeserializeObject<dynamic>(await response.Content.ReadAsStringAsync());
            Assert.IsNotNull(content);
            Assert.IsNotNull(content.data);
            Assert.IsInstanceOf<JObject>(content.data);
            Assert.That(IsAmenityObjEqualToAmenity(content.data, amenity));
        }

        [Test]
        public async Task GetAmenityById_Two()
        {
            var amenity1 = AddAmenity(NewAmenity());
            var amenity2 = AddAmenity(NewAmenity(name: "name", description: "desc"));
            var response1 = await _client.GetAsync($"/amenities/{amenity1.Id}");
            var content1 = JsonConvert.DeserializeObject<dynamic>(await response1.Content.ReadAsStringAsync());
            Assert.IsNotNull(content1);
            Assert.IsNotNull(content1.data);
            Assert.IsInstanceOf<JObject>(content1.data);
            Assert.That(IsAmenityObjEqualToAmenity(content1.data, amenity1));

            var response2 = await _client.GetAsync($"/amenities/{amenity2.Id}");
            var content2 = JsonConvert.DeserializeObject<dynamic>(await response2.Content.ReadAsStringAsync());
            Assert.IsNotNull(content2);
            Assert.IsNotNull(content2.data);
            Assert.IsInstanceOf<JObject>(content2.data);
            Assert.That(IsAmenityObjEqualToAmenity(content2.data, amenity2));
        }

        [Test]
        public async Task DeleteAmenities_Empty()
        {
            var response = await _client.DeleteAsync("/amenities/1");
            Assert.That(response.StatusCode, Is.EqualTo(System.Net.HttpStatusCode.NotFound));
        }

        [Test]
        public async Task DeleteAmenities_OneNotFound()
        {
            var amenity = AddAmenity(NewAmenity());
            var response = await _client.DeleteAsync($"/amenities/{amenity.Id+1}");
            Assert.That(response.StatusCode, Is.EqualTo(System.Net.HttpStatusCode.NotFound));
        }

        [Test]
        public async Task DeleteAmenities_OneFound()
        {
            var amenity = AddAmenity(NewAmenity());
            Assert.That(DoesAmenityExist(amenity.Id), Is.True);

            var response = await _client.DeleteAsync($"/amenities/{amenity.Id}");
            response.EnsureSuccessStatusCode();
            Assert.That(DoesAmenityExist(amenity.Id), Is.False);
        }

        [Test]
        public async Task DeleteAmenities_Two()
        {
            var amenity1 = AddAmenity(NewAmenity());
            var amenity2 = AddAmenity(NewAmenity(name: "name", description: "desc"));

            Assert.That(DoesAmenityExist(amenity1.Id), Is.True);
            Assert.That(DoesAmenityExist(amenity2.Id), Is.True);

            var response = await _client.DeleteAsync($"/amenities/{amenity1.Id}");
            response.EnsureSuccessStatusCode();
            Assert.That(DoesAmenityExist(amenity1.Id), Is.False);
            Assert.That(DoesAmenityExist(amenity2.Id), Is.True);
        }

        [Test]
        public async Task AddAmenities_Ok()
        {
            var amenity = NewAmenity();
            using StringContent jsonContent = new(
                System.Text.Json.JsonSerializer.Serialize(new
                {  
                    id = 1,
                    name = amenity.Name,
                    description = amenity.Description,
                    startTime = amenity.StartTime.ToString(),
                    endTime = amenity.EndTime.ToString()
                }),
                Encoding.UTF8,
                "application/json");
            var response = await _client.PostAsync("/amenities", jsonContent);
            response.EnsureSuccessStatusCode();
            var content = JsonConvert.DeserializeObject<dynamic>(await response.Content.ReadAsStringAsync());
            Assert.IsNotNull(content);
            Assert.IsNotNull(content.data);
            Assert.IsInstanceOf<JObject>(content.data);
            Assert.IsNotNull(content.data.id);
            amenity.Id = content.data.id;

            Assert.That(DoesAmenityExist(amenity.Id), Is.True);
            Assert.That(IsAmenityObjEqualToAmenity(content.data, amenity), Is.True);
        }

        [Test]
        public async Task AddAmenities_NoName()
        {
            var amenity = NewAmenity();
            using StringContent jsonContent = new(
                System.Text.Json.JsonSerializer.Serialize(new
                {
                    description = amenity.Description,
                    startTime = amenity.StartTime.ToString(),
                    endTime = amenity.EndTime.ToString()
                }),
                Encoding.UTF8,
                "application/json");
            var response = await _client.PostAsync("/amenities", jsonContent);
            Assert.That(response.StatusCode, Is.EqualTo(System.Net.HttpStatusCode.BadRequest));
        }

        [Test]
        public async Task AddAmenities_NoDesc()
        {
            var amenity = NewAmenity();
            using StringContent jsonContent = new(
                System.Text.Json.JsonSerializer.Serialize(new
                {
                    name = amenity.Name,
                    startTime = amenity.StartTime.ToString(),
                    endTime = amenity.EndTime.ToString()
                }),
                Encoding.UTF8,
                "application/json");
            var response = await _client.PostAsync("/amenities", jsonContent);
            Assert.That(response.StatusCode, Is.EqualTo(System.Net.HttpStatusCode.BadRequest));
        }

        [Test]
        public async Task AddAmenities_NoStart()
        {
            var amenity = NewAmenity();
            using StringContent jsonContent = new(
                System.Text.Json.JsonSerializer.Serialize(new
                {
                    name = amenity.Name,
                    description = amenity.Description,
                    endTime = amenity.EndTime.ToString()
                }),
                Encoding.UTF8,
                "application/json");
            var response = await _client.PostAsync("/amenities", jsonContent);
            response.EnsureSuccessStatusCode();
            var content = JsonConvert.DeserializeObject<dynamic>(await response.Content.ReadAsStringAsync());
            Assert.IsNotNull(content);
            Assert.IsNotNull(content.data);
            Assert.IsInstanceOf<JObject>(content.data);
            Assert.IsNotNull(content.data.id);
            amenity.Id = content.data.id;
            amenity.StartTime = new TimeSpan(0, 0, 0);

            Assert.That(DoesAmenityExist(amenity.Id), Is.True);
            Assert.That(IsAmenityObjEqualToAmenity(content.data, amenity), Is.True);
        }

        [Test]
        public async Task AddAmenities_NoEnd()
        {
            var amenity = NewAmenity();
            using StringContent jsonContent = new(
                System.Text.Json.JsonSerializer.Serialize(new
                {
                    name = amenity.Name,
                    description = amenity.Description,
                    startTime = amenity.StartTime.ToString()
                }),
                Encoding.UTF8,
                "application/json");
            var response = await _client.PostAsync("/amenities", jsonContent);
            Assert.That(response.StatusCode, Is.EqualTo(System.Net.HttpStatusCode.BadRequest));
        }

        [Test]
        public async Task PutAmenities_NotFound()
        {
            var amenity = NewAmenity();
            using StringContent jsonContent = new(
                System.Text.Json.JsonSerializer.Serialize(new
                {
                    name = amenity.Name,
                    description = amenity.Description,
                    startTime = amenity.StartTime.ToString(),
                    endTime = amenity.EndTime.ToString()
                }),
                Encoding.UTF8,
                "application/json");
            var response = await _client.PutAsync("/amenities/1", jsonContent);
            Assert.That(response.StatusCode, Is.EqualTo(System.Net.HttpStatusCode.NotFound));
        }

        [Test]
        public async Task PutAmenities_ok()
        {
            var amenity = AddAmenity(NewAmenity());

            amenity.Name = "Other";
            amenity.Description = "OtherDesc";
            amenity.StartTime = new TimeSpan(05, 11, 11);
            amenity.EndTime = new TimeSpan(06, 11, 11);

            using StringContent jsonContent = new(
                System.Text.Json.JsonSerializer.Serialize(new
                {
                    name = amenity.Name,
                    description = amenity.Description,
                    startTime = amenity.StartTime.ToString(),
                    endTime = amenity.EndTime.ToString()
                }),
                Encoding.UTF8,
                "application/json");
            var response = await _client.PutAsync($"/amenities/{amenity.Id}", jsonContent);
            response.EnsureSuccessStatusCode();
            var content = JsonConvert.DeserializeObject<dynamic>(await response.Content.ReadAsStringAsync());
            Assert.IsNotNull(content);
            Assert.IsNotNull(content.data);
            Assert.IsInstanceOf<JObject>(content.data);
            Assert.That(IsAmenityObjEqualToAmenity(content.data, amenity), Is.True);
        }

        [Test]
        public async Task PutAmenities_NoName()
        {
            var amenity = AddAmenity(NewAmenity());
            var newAmenity = new Amenity("Other", "otherDesc", new TimeSpan(05, 11, 11), new TimeSpan(06, 11, 11));

            using StringContent jsonContent = new(
                System.Text.Json.JsonSerializer.Serialize(new
                {
                    description = newAmenity.Description,
                    startTime = newAmenity.StartTime.ToString(),
                    endTime = newAmenity.EndTime.ToString()
                }),
                Encoding.UTF8,
                "application/json");
            var response = await _client.PutAsync($"/amenities/{amenity.Id}", jsonContent);
            Assert.That(response.StatusCode, Is.EqualTo(System.Net.HttpStatusCode.BadRequest));

            var foundAmenity = GetAmenity(amenity.Id);
            Assert.That(AreAmenitiesEqual(amenity, foundAmenity!), Is.True); //Ensure amenity wasn't changed
        }

        [Test]
        public async Task PutAmenities_NoDesc()
        {
            var amenity = AddAmenity(NewAmenity());
            var newAmenity = new Amenity("Other", "otherDesc", new TimeSpan(05, 11, 11), new TimeSpan(06, 11, 11));

            using StringContent jsonContent = new(
                System.Text.Json.JsonSerializer.Serialize(new
                {
                    name = newAmenity.Name,
                    startTime = newAmenity.StartTime.ToString(),
                    endTime = newAmenity.EndTime.ToString()
                }),
                Encoding.UTF8,
                "application/json");
            var response = await _client.PutAsync($"/amenities/{amenity.Id}", jsonContent); ;
            Assert.That(response.StatusCode, Is.EqualTo(System.Net.HttpStatusCode.BadRequest));

            var foundAmenity = GetAmenity(amenity.Id);
            Assert.That(AreAmenitiesEqual(amenity, foundAmenity!), Is.True); //Ensure amenity wasn't changed
        }

        [Test]
        public async Task PutAmenities_NoStart()
        {
            var amenity = AddAmenity(NewAmenity());
            var newAmenity = new Amenity(amenity.Id, "Other", "otherDesc", new TimeSpan(05, 11, 11), new TimeSpan(06, 11, 11));

            using StringContent jsonContent = new(
                System.Text.Json.JsonSerializer.Serialize(new
                {
                    name = newAmenity.Name,
                    description = newAmenity.Description,
                    endTime = newAmenity.EndTime.ToString()
                }),
                Encoding.UTF8,
                "application/json");
            var response = await _client.PutAsync($"/amenities/{amenity.Id}", jsonContent);
            response.EnsureSuccessStatusCode();
            var content = JsonConvert.DeserializeObject<dynamic>(await response.Content.ReadAsStringAsync());
            newAmenity.StartTime = new TimeSpan(0, 0, 0);
            Assert.IsNotNull(content);
            Assert.IsNotNull(content.data);
            Assert.IsInstanceOf<JObject>(content.data);
            Assert.That(IsAmenityObjEqualToAmenity(content.data, newAmenity), Is.True);
        }

        [Test]
        public async Task PutAmenities_NoEnd()
        {
            var amenity = AddAmenity(NewAmenity());
            var newAmenity = new Amenity("Other", "otherDesc", new TimeSpan(05, 11, 11), new TimeSpan(06, 11, 11));

            using StringContent jsonContent = new(
                System.Text.Json.JsonSerializer.Serialize(new
                {
                    name = newAmenity.Name,
                    description = newAmenity.Description,
                    startTime = newAmenity.StartTime.ToString()
                }),
                Encoding.UTF8,
                "application/json");
            var response = await _client.PutAsync($"/amenities/{amenity.Id}", jsonContent);
            Assert.That(response.StatusCode, Is.EqualTo(System.Net.HttpStatusCode.BadRequest));

            var foundAmenity = GetAmenity(amenity.Id);
            Assert.That(AreAmenitiesEqual(amenity, foundAmenity!), Is.True); //Ensure amenity wasn't changed
        }


        [TearDown]
        public void TearDown() 
        {
            _client.Dispose();
            _appFactory.Dispose();
        }

        private Amenity AddAmenity(Amenity amenity)
        {
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

        private bool DoesAmenityExist(int id)
        {
            Amenity? amenity = null;
            using var connection = new NpgsqlConnection(_connectionString);
            connection.Open();
            using var command = new NpgsqlCommand("SELECT id, name, description, start_time, end_time FROM amenities where id = @id;", connection);
            command.Parameters.AddWithValue("id", id);
            using var reader = command.ExecuteReader();
            if (reader.Read())
            {
                amenity = new Amenity(
                   reader.GetInt32(0),
                   reader.GetString(1),
                   reader.GetString(2),
                   reader.GetTimeSpan(3),
                   reader.GetTimeSpan(4)
                );
            }

            return amenity!=null;
        }

        private Amenity? GetAmenity(int id)
        {
            Amenity? amenity = null;
            using var connection = new NpgsqlConnection(_connectionString);
            connection.Open();
            using var command = new NpgsqlCommand("SELECT id, name, description, start_time, end_time FROM amenities where id = @id;", connection);
            command.Parameters.AddWithValue("id", id);
            using var reader = command.ExecuteReader();
            if (reader.Read())
            {
                amenity = new Amenity(
                   reader.GetInt32(0),
                   reader.GetString(1),
                   reader.GetString(2),
                   reader.GetTimeSpan(3),
                   reader.GetTimeSpan(4)
                );
            }

            return amenity;
        }
        private static Amenity NewAmenity(string name = "test name", string description = "test description", TimeSpan? start = null, TimeSpan? end = null)
        {
            return new Amenity(0, name, description, start ?? new TimeSpan(1, 0, 0), end ?? new TimeSpan(2, 0, 0));
        }

        private static bool IsAmenityObjEqualToAmenity(dynamic a, Amenity b)
        {
            var i1 = a.id == b.Id; 
            var i2 = a.name == b.Name;
            var i3 = a.description == b.Description;
            var i4 = a.startTime.ToString().Equals(b.StartTime.ToString());
            var i5 = a.endTime.ToString().Equals(b.EndTime.ToString());
            return i1 && i2 && i3 && i4 && i5;
        }

        private static bool AreAmenitiesEqual(Amenity a, Amenity b)
        {
            return a.Id == b.Id && a.Name == b.Name && a.Description == b.Description && a.StartTime.Equals(b.StartTime) && a.EndTime.Equals(b.EndTime);
        }
    }
}
