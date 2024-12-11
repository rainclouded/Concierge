import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { fetchWithAuth } from "../utils/authFetch";
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
          `/amenities/`
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
    <div className="relative bg-primary min-h-screen pb-10">
      <Link
        to="/home"
        className="fixed left-4 top-4 p-1 text-xl font-semibold rounded-full z-50"
      >
        <FontAwesomeIcon icon={faArrowLeft} />
      </Link>
      {/* Sticky Header */}
      <header className="sticky bg-white top-0 p-3 flex justify-center items-center z-40">
        <h1 className="text-2xl mt-2 font-extrabold mx-auto">Amenities</h1>
      </header>

      {error && <p className="text-red-700">Error: {error}</p>}

      {!error && amenities.length > 0 ? (
        <div className="flex flex-col justify-center items-center space-y-4 mt-4">
          <p className="text-sm italic text-justify px-4 mb-2 text-lightText lg:max-w-[50%]">
            Explore our range of amenities designed to make your stay
            unforgettable. From luxurious facilities to convenient services, we
            offer everything you need for a relaxing and enjoyable experience.
            Please take advantage of these offerings and make the most of your
            time with us!
          </p>
          {amenities.map((amenity) => (
            <div
              key={amenity.id}
              className="flex rounded-lg p-4 shadow-md bg-lightPrimary w-full max-w-[95%] lg:max-w-[50%]"
            >
              <div className="basis-1/2 flex items-center">
                <h2 className="text-3xl font-bold text-black">
                  {amenity.name}
                </h2>
              </div>
              <div className="basis-1/2">
                <p className="text-brown mb-2">{amenity.description}</p>
                <div className="text-brown text-sm font-medium">
                  <p>
                    Opens:{" "}
                    {new Date(`1970-01-01T${amenity.startTime}`)
                      .toLocaleTimeString([], {
                        hour: "numeric",
                        minute: "2-digit",
                        hour12: true,
                      })
                      .toLowerCase()}
                  </p>
                  <p>
                    Closes:{" "}
                    {new Date(`1970-01-01T${amenity.endTime}`)
                      .toLocaleTimeString([], {
                        hour: "numeric",
                        minute: "2-digit",
                        hour12: true,
                      })
                      .toLowerCase()}
                  </p>
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
