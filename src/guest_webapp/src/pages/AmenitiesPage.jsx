import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { fetchWithAuth } from "../utils/authFetch"

const AmenitiesPage = () => {
  const [amenities, setAmenities] = useState([]);
  const [error, setError] = useState(null);

  // Fetch amenities from the backend
  useEffect(() => {
    const fetchAmenities = async () => {
      try {
        const response = await fetchWithAuth(
          `${import.meta.env.VITE_API_BASE_URL}/amenities/`
        );
        if (!response.ok) {
          throw new Error("Failed to fetch amenities");
        }
        const data = await response.json();
        setAmenities(data.data);
      } catch (err) {
        setError(err.message);
      }
    };
    fetchAmenities();
  }, []);

  return (
    <div className="p-6 text-center">
      <h1 className="text-2xl font-bold mb-4">Amenities</h1>

      {error && <p className="text-red-500">Error: {error}</p>}

      {!error && amenities.length > 0 ? (
        <table className="table-auto w-full border-collapse border border-gray-300">
          <thead>
            <tr>
              <th className="border border-gray-300 px-4 py-2">Name</th>
              <th className="border border-gray-300 px-4 py-2">Description</th>
              <th className="border border-gray-300 px-4 py-2">Start Time</th>
              <th className="border border-gray-300 px-4 py-2">End Time</th>
            </tr>
          </thead>
          <tbody>
            {amenities.map((amenity) => (
              <tr key={amenity.id} class="amenity-row-item">
                <td className="border border-gray-300 px-4 py-2">
                  {amenity.name}
                </td>
                <td className="border border-gray-300 px-4 py-2">
                  {amenity.description}
                </td>
                <td className="border border-gray-300 px-4 py-2">
                  {amenity.startTime}
                </td>
                <td className="border border-gray-300 px-4 py-2">
                  {amenity.endTime}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        !error && <p>Loading amenities...</p>
      )}

      {/* Back to Homepage Button */}
      <Link
        to="/home"
        className="inline-block mb-4 mt-4 px-6 py-3 bg-blue-500 text-white rounded"
      >
        Back to Homepage
      </Link>
    </div>
  );
};

export default AmenitiesPage;

