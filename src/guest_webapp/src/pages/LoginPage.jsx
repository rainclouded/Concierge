import React, { useState } from "react";

const LoginPage = () => {
  const [roomKey, setRoomKey] = useState("");

  const handleSubmit = (event) => {
    event.preventDefault();
    if (roomKey) {
      window.location.href = "/home";
    } else {
      alert("Invalid room key");
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gradient-to-r from-beige-200 to-beige-500 p-4">
      <div className="bg-white p-6 sm:p-8 lg:p-10 rounded-lg shadow-xl max-w-md w-full">
        <h1 className="text-2xl sm:text-3xl font-semibold text-center text-gray-800 mb-4 sm:mb-6">
          Welcome
        </h1>
        <form onSubmit={handleSubmit}>
          <label
            htmlFor="roomKey"
            className="block text-sm font-medium text-gray-700 mb-2"
          >
            Enter Your Room Key:
          </label>
          <input
            type="text"
            id="roomKey"
            value={roomKey}
            onChange={(e) => setRoomKey(e.target.value)}
            className="block w-full p-2 sm:p-3 border border-gray-300 rounded-md mb-3 sm:mb-4 focus:ring-2 focus:ring-blue-500 shadow-sm hover:shadow-md transition duration-200"
            placeholder="Room Key"
          />
          <button
            type="submit"
            className="w-full bg-blue-600 text-white py-2 sm:py-3 rounded-md hover:bg-blue-700 transition duration-200"
          >
            Submit
          </button>
        </form>
      </div>
    </div>
  );
};

export default LoginPage;

