import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { fetchWithAuth } from "../utils/authFetch";
import { Link } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faArrowLeft } from "@fortawesome/free-solid-svg-icons";

const IncidentReportPage = () => {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [message, setMessage] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    const requestBody = {
      title: title,
      description: description,
      filing_person_id: 1234, // Replace with actual person ID
      reviewer_id: 5678, // Add reviewer_id, replace with actual logic
      severity: "LOW", // Default value for severity
      status: "OPEN", // Default value for status
    };

    try {
      const response = await fetchWithAuth(
        `${import.meta.env.VITE_API_BASE_URL}/incident_reports/`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(requestBody),
        }
      );
      if (!response.ok) throw new Error("Failed to submit report");
      setMessage("Incident report submitted successfully");
      setTitle("");
      setDescription("");
    } catch (error) {
      setMessage("Error: " + error.message);
    }
  };

  return (
    <div className="relative pt-1 bg-primary h-screen">
      <Link to="/home" className="absolute left-4 top-4 p-1 text-xl font-semibold rounded-full z-50">
        <FontAwesomeIcon icon={faArrowLeft} />
      </Link>
      {/* Sticky Header */}
      <header className="sticky top-0 my-2 p-2 flex justify-center items-center z-40">
        <h1 className="text-2xl font-extrabold mx-auto">Make a report</h1>
      </header>
      <div className="w-full max-w-[95%] lg:max-w-[50%] mx-auto">
        {message && <p className="text-center">{message}</p>}
        <form onSubmit={handleSubmit}>
  <div className="mb-4">
    <label className="block text-brown font-semibold mb-1" htmlFor="title">
      Title
    </label>
    <input
      type="text"
      id="title"
      className="p-2 border border-gray-300 rounded-lg w-full"
      placeholder="Title"
      value={title}
      onChange={(e) => setTitle(e.target.value)}
    />
  </div>

  <div className="mb-4">
    <label className="block text-brown font-semibold mb-1" htmlFor="description">
      Description
    </label>
    <textarea
      id="description"
      className="p-2 border border-gray-300 rounded-lg h-48 w-full"
      placeholder="Describe the incident..."
      value={description}
      onChange={(e) => setDescription(e.target.value)}
    />
  </div>

  <div className="flex justify-end">
    <button
      type="submit"
      className="px-4 py-2 bg-black text-white rounded-lg mr-2"
    >
      Submit
    </button>
    <button
      type="button"
      className="px-4 py-2 bg-red-700 text-white rounded-lg"
      onClick={() => navigate(-1)}
    >
      Cancel
    </button>
  </div>
</form>

      </div>
    </div>
  );
};

export default IncidentReportPage;
