import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { fetchWithAuth } from "../utils/authFetch";
import { Link } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faArrowLeft } from "@fortawesome/free-solid-svg-icons";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/ReactToastify.css";
import { getAccountId } from "../utils/auth";

const IncidentReportPage = () => {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [message, setMessage] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    const accountId = getAccountId();
    const filingPersonId = accountId === 0 ? 1 : accountId; // Retrieve accountId from token
    if (!filingPersonId && filingPersonId !== 0) {
      toast.error("Unable to submit. No account ID found.");
      return;
    }

    const requestBody = {
      title: title,
      description: description,
      filing_person_id: filingPersonId,
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
      toast.success("Incident report submitted successfully!");
      setTitle("");
      setDescription("");
    } catch (error) {
      toast.error(`Error: ${error.message}`);
    }
  };

  return (
    <div className="relative bg-primary min-h-screen pb-10">
      <ToastContainer bodyClassName="toast-body" />
      <Link
        to="/home"
        className="fixed left-4 top-4 p-1 text-xl font-semibold rounded-full z-50"
      >
        <FontAwesomeIcon icon={faArrowLeft} />
      </Link>
      {/* Sticky Header */}
      <header className="sticky bg-white top-0 p-3 flex justify-center items-center z-40">
        <h1 className="text-2xl mt-2 font-extrabold mx-auto">Make a report</h1>
      </header>
      <div className="w-full max-w-[95%] lg:max-w-[50%] mx-auto">
        <form onSubmit={handleSubmit} className="px-2">
          <div className="mb-4">
            <p className="text-sm italic text-justify lg:px-0 my-4 text-lightText px-1">
              Have a concern? File a complaint an issue quickly and
              easily. Provide title and detailed description and our team
              will resolve it promptly. Thank you for your patience!
            </p>
            <label
              className="block text-brown font-semibold mb-1"
              htmlFor="title"
            >
              Title
            </label>
            <p className="text-sm italic text-justify lg:px-0 mb-2 text-lightText px-1">
              eg. Strange noise
            </p>
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
            <label
              className="block text-brown font-semibold mb-1"
              htmlFor="description"
            >
              Description
            </label>
            <p className="text-sm italic text-justify lg:px-0 mb-2 text-lightText px-1">
              eg. Strange noise coming from the wall to the left of the bathroom door. Please handle the issue asap. Thank you.
            </p>
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
