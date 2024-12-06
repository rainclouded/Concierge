using amenities_server.model;
using Npgsql;

namespace amenities_server.persistence
{
    public class PostgresAmenityPersistence : IAmenityPersistence
    {
        private readonly string _connectionString;

        /*
        Constructor that initializes the PostgresAmenityPersistence with a 
            connection string and validates the database connection.
        Args:
            connectionString: The connection string used to connect to the PostgreSQL database.
        */
        public PostgresAmenityPersistence(string connectionString)
        {
            _connectionString = connectionString;
            using var connection = new NpgsqlConnection(_connectionString);
            try
            {
                connection.Open(); //validate connection
            }
            catch (Exception ex)
            {
                throw new InvalidOperationException("Could not establish a connection to the database.", ex);
            }
        }

        /*
        Adds a new Amenity to the database and returns the added 
        Amenity with the generated ID.
        Args:
            amenity: The Amenity object that contains the data to be added to the database.
        Returns:
            Amenity: The added Amenity with the assigned ID from the database.
        */
        public Amenity AddAmenity(Amenity amenity)
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

        /*
        Deletes an Amenity from the database based on its ID.
        Args:
            id: The ID of the Amenity to be deleted.
        */
        public void DeleteAmenity(int id)
        {
            using var connection = new NpgsqlConnection(_connectionString);
            connection.Open();
            using var command = new NpgsqlCommand("DELETE FROM amenities where id = @id;", connection);
            command.Parameters.AddWithValue("id", id);
            using var reader = command.ExecuteReader();
        }

        /*
        Retrieves all Amenities from the database.
        Returns:
            IEnumerable<Amenity>: A collection of Amenity objects.
        */
        public IEnumerable<Amenity> GetAmenities()
        {
            var amenities = new List<Amenity>();
            using var connection = new NpgsqlConnection(_connectionString);
            connection.Open();
            using var command = new NpgsqlCommand("SELECT id, name, description, start_time, end_time FROM amenities;", connection);
            using var reader = command.ExecuteReader();
            while (reader.Read())
            {
                amenities.Add(new Amenity(
                   reader.GetInt32(0),
                   reader.GetString(1),
                   reader.GetString(2),
                   reader.GetTimeSpan(3),
                   reader.GetTimeSpan(4)
                ));
            }
            return amenities;
        }

        /*
        Retrieves a specific Amenity from the database by its ID.
        Args:
            id: The ID of the Amenity to retrieve.
        Returns:
            Amenity: The Amenity object matching the provided ID, or null if not found.
        */
        public Amenity GetAmenityByID(int id)
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

        /*
        Updates an existing Amenity in the database with new values.
        Args:
            id: The ID of the Amenity to be updated.
            amenity: The Amenity object containing the updated values.
        Returns:
            Amenity: The updated Amenity object.
        */
        public Amenity UpdateAmenity(int id, Amenity amenity)
        {
            using var connection = new NpgsqlConnection(_connectionString);
            connection.Open();
            using var command = new NpgsqlCommand("UPDATE amenities SET name=@name, description=@description, start_time=@startTime, end_time=@endtime where id = @id;", connection);
            command.Parameters.AddWithValue("id", id);
            command.Parameters.AddWithValue("name", amenity.Name);
            command.Parameters.AddWithValue("description", amenity.Description);
            command.Parameters.AddWithValue("startTime", amenity.StartTime);
            command.Parameters.AddWithValue("endTime", amenity.EndTime);
            using var reader = command.ExecuteReader();

            return GetAmenityByID(id);
        }
    }
}
