import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const IncidentReportPage = () => {
  const [incidentText, setIncidentText] = useState("");
  const [message, setMessage] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch(`${import.meta.env.VITE_API_BASE_URL}/incident_reports/`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ incident: incidentText }),
      });
      if (!response.ok) throw new Error("Failed to submit report");
      setMessage("Incident report submitted successfully");
      setIncidentText("");
    } catch (error) {
      setMessage("Error: " + error.message);
    }
  };

  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-4">Report an Incident</h1>
      {message && <p>{message}</p>}
      <form onSubmit={handleSubmit}>
        <textarea
          className="p-2 border border-gray-300 rounded mb-4 h-48"
          placeholder="Describe the incident..."
          value={incidentText}
          onChange={(e) => setIncidentText(e.target.value)}
        />
        <div className="flex justify-end">
          <button
            type="button"
            className="px-4 py-2 bg-gray-500 text-white rounded mr-2"
            onClick={() => navigate(-1)}  // Navigate back to the previous page
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
  );
};

export default IncidentReportPage;
