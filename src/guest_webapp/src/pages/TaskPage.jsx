import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { fetchWithAuth } from "../utils/authFetch";
import { faArrowLeft } from "@fortawesome/free-solid-svg-icons";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/ReactToastify.css";

function TaskPage() {
  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    let isMounted = true;
    const fetchTasks = async () => {
      try {
        const response = await fetchWithAuth(
          `${import.meta.env.VITE_API_BASE_URL}/tasks/`,
          {
            method: "GET",
            headers: {
              "Content-Type": "application/json",
            },
          }
        );
        if (!response.ok) {
          throw new Error(`Error: ${response.status} ${response.statusText}`);
        }
        const data = await response.json();
        if (isMounted) {
          setTasks(data.data);
          console.log(data.data);
        }
      } catch (err) {
        if (isMounted) {
          toast(`Failed to load your requests: ${err.message}`);
        }
      }
    };
    fetchTasks();
    return () => {
      isMounted = false;
    };
  }, []);

  function parseCamelCase(input) {
    return input
      .replace(/([a-zA-Z])([0-9])/g, "$1 $2")
      .replace(/([a-z])([A-Z])/g, "$1 $2");
  }

  function formatDate(dateString) {
    const date = new Date(dateString);

    const months = [
      "Jan",
      "Feb",
      "Mar",
      "Apr",
      "May",
      "Jun",
      "Jul",
      "Aug",
      "Sep",
      "Oct",
      "Nov",
      "Dec",
    ];
    const month = months[date.getUTCMonth()];
    const day = date.getUTCDate();

    // Add ordinal suffix
    const dayWithSuffix =
      day +
      (day % 10 === 1 && day !== 11
        ? "st"
        : day % 10 === 2 && day !== 12
        ? "nd"
        : day % 10 === 3 && day !== 13
        ? "rd"
        : "th");

    const year = date.getUTCFullYear();

    let hours = date.getUTCHours();
    const minutes = date.getUTCMinutes();
    const amPm = hours >= 12 ? "pm" : "am";
    hours = hours % 12 || 12; // Convert to 12-hour format

    return `${month}, ${dayWithSuffix} ${year} at ${hours}:${minutes
      .toString()
      .padStart(2, "0")}${amPm}`;
  }

  function getStatusBackground(status) {
    switch (status) {
      case "Pending":
        return "bg-lightRed";
      case "Completed":
        return "bg-lightGreen";
      case "In Progress":
        return "bg-yellow-300";
      default:
        return "bg-primary";
    }
  }

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
        <h1 className="text-2xl mt-2 font-extrabold mx-auto">Your requests</h1>
      </header>

      <div className="flex flex-col justify-center items-center space-y-4 mt-4">
        <p className="text-sm italic text-justify px-4 text-lightText lg:max-w-[50%]">
          Here are your most recent requests. Stay on top of your progress and
          keep track of all updates in one place. We've made it easy for you to
          monitor the status of your requests and ensure everything is moving
          forward smoothly.
        </p>
        <div className="lg:p-4 w-full max-w-[95%] lg:max-w-[50%]">
          {tasks.length === 0 ? (
            <p>No requests made.</p>
          ) : (
            <ul className="flex flex-col space-y-4">
              {tasks.map((task) => (
                <li key={task.id} className="relative min-h-36 bg-lightPrimary px-3 py-2 rounded-lg flex flex-col shadow-md">
                  <div className="text-xl font-semibold tracking-wide">{parseCamelCase(task.taskType)}</div>
                  <div className="italic text-sm text-lightText">Created on {formatDate(task.createdAt)}</div>
                  <div className="text-brown">{task.description}</div>
                  <div className="mt-auto">
                    <span className={`px-2 py-0.5 text-white rounded-lg ${getStatusBackground(
                      task.status
                    )}`}>
                      {task.status}
                    </span>
                  </div>
                </li>
              ))}
            </ul>
          )}
        </div>
      </div>
    </div>
  );
}

export default TaskPage;
