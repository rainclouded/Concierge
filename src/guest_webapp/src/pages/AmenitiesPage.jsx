import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { fetchWithAuth } from "../utils/authFetch"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faArrowLeft } from "@fortawesome/free-solid-svg-icons";

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
    <div className="relative pt-1 bg-primary h-screen">
      <Link to="/home" className="absolute left-4 top-4 p-1 text-xl font-semibold rounded-full z-50">
        <FontAwesomeIcon icon={faArrowLeft} />
      </Link>
      {/* Sticky Header */}
      <header className="sticky top-0 my-2 p-2 flex justify-center items-center z-40">
        <h1 className="text-2xl font-extrabold mx-auto">Amenities</h1>
      </header>

      {error && <p className="text-red-700">Error: {error}</p>}

      {!error && amenities.length > 0 ? (
        <div className="flex flex-col justify-center items-center space-y-4 mt-4">
          {amenities.map((amenity) => (
            <div
              key={amenity.id}
              className="flex rounded-lg p-4 shadow-md bg-lightPrimary w-full max-w-[95%] lg:max-w-[50%]"
            >
              <div className="basis-1/2 flex items-center">
                <h2 className="text-3xl font-bold text-black">{amenity.name}</h2>
              </div>
              <div className="basis-1/2">
              <p className="text-brown mb-2">{amenity.description}</p>
                <div className="text-brown text-sm font-medium">
                  <p>Start Time: {amenity.startTime}</p>
                  <p>End Time: {amenity.endTime}</p>
                </div>
              </div>
            </div>
          ))}
        </div>
      ) : (
        !error && <p>Loading amenities...</p>
      )}
    </div>
  );
};

export default AmenitiesPage;

