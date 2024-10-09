import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

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
      filing_person_id: 1234,  // Replace with actual person ID
      reviewer_id: 5678,       // Add reviewer_id, replace with actual logic
      severity: "LOW",         // Default value for severity
      status: "OPEN",          // Default value for status
    };

    try {
      const response = await fetch(`${import.meta.env.VITE_API_BASE_URL}/incident_reports/`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(requestBody),
      });
      if (!response.ok) throw new Error("Failed to submit report");
      setMessage("Incident report submitted successfully");
      setTitle("");
      setDescription("");
    } catch (error) {
      setMessage("Error: " + error.message);
    }
  };

  return (
    <div className="p-6 flex justify-center">
      <div className="w-full max-w-full md:max-w-[50%] mx-auto">
        <h1 className="text-2xl font-bold mb-4 text-center">Report an Incident</h1>
        {message && <p className="text-center">{message}</p>}
        <form onSubmit={handleSubmit}>
          <input
            type="text"
            className="p-2 border border-gray-300 rounded mb-4 w-full"
            placeholder="Title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
          />
          <textarea
            className="p-2 border border-gray-300 rounded mb-4 h-48 w-full"
            placeholder="Describe the incident..."
            value={description}
            onChange={(e) => setDescription(e.target.value)}
          />
          <div className="flex justify-end">
            <button
              type="button"
              className="px-4 py-2 bg-gray-500 text-white rounded mr-2"
              onClick={() => navigate(-1)}
            >
              Back
            </button>

            <button
              type="submit"
              className="px-4 py-2 bg-blue-500 text-white rounded"
            >
              Submit
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default IncidentReportPage;
