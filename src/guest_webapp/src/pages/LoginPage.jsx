import React, { useState } from "react";
import { setSessionKey } from "../utils/auth";
import conciergeVector from "../assets/conciergevector.svg";

const LoginPage = () => {
  const [roomKey, setRoomKey] = useState("");
  const [roomNum, setRoomNum] = useState("");
  const [error, setError] = useState(null);

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      const response = await fetch(
        `/sessions`,
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            username: `${roomNum}`,
            password: `${roomKey}`,
          }),
        }
      );
      if (!response.ok) {
        const errMsg = await response.json();
        setError(errMsg.error || errMsg.message || "Login Failed");
        return;
      }
      const data = await response.json();
      setSessionKey(data.data.sessionKey);
      sessionStorage.setItem("roomKey", roomKey);
      sessionStorage.setItem("roomNum", roomNum);
      window.location.href = "/home";
    } catch (error) {
      console.log(error);
      setError("Login Failed");
    }
  };

  return (
    <div className="flex bg-secondary items-center justify-center min-h-screen bg-gradient-to-r from-beige-200 to-beige-500 p-4">
      <div className="bg-lightPrimary p-6 sm:p-8 lg:p-10 rounded-lg shadow-xl max-w-md w-full">
        <h1 className="text-2xl sm:text-3xl font-semibold text-center text-black">
          Welcome to
        </h1>
        <div className="flex justify-center mb-5">
          <img
            src={conciergeVector}
            alt="Concierge"
            className="h-24 object-contain"
          />
        </div>
        <form onSubmit={handleSubmit}>
          <label
            htmlFor="roomNum"
            className="block text-sm font-medium text-gray-700 mb-2"
          >
            Enter Your Key Number:
          </label>
          <input
            type="text"
            id="roomNum"
            value={roomNum}
            onChange={(e) => setRoomNum(e.target.value)}
            className="block w-full p-2 sm:p-3 border border-gray-300 rounded-md mb-3 sm:mb-4 focus:ring-2 focus:ring-primary shadow-sm hover:shadow-md transition duration-200"
            placeholder="Room #"
          />
          
          <label
            htmlFor="roomKey"
            className="block text-sm font-medium text-gray-700 mb-2"
          >
            Enter Your Key Code:
          </label>
          <input
            type="text"
            id="roomKey"
            value={roomKey}
            onChange={(e) => setRoomKey(e.target.value)}
            className="block w-full p-2 sm:p-3 border border-gray-300 rounded-md mb-3 sm:mb-4 focus:ring-2 focus:ring-primary shadow-sm hover:shadow-md transition duration-200"
            placeholder="Key Code"
          />
          {error && <p className="text-red-500">Error: {error}</p>}
          <button
            type="submit"
            className="w-full bg-black text-white py-2 sm:py-3 rounded-md hover:bg-primary hover:text-black transition duration-200"
          >
            Login
          </button>
        </form>
      </div>
    </div>
  );
};

export default LoginPage;
