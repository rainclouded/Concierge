import React, { useState } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faConciergeBell,
  faBars,
  faUser,
  faExclamationTriangle,
  faBroom,
  faClock,
  faHamburger,
  faShirt,
  faSpa,
  faWrench,
  faCog,
  faSignOutAlt,
  faCalendar,
  faClipboardList,
} from "@fortawesome/free-solid-svg-icons";

import ServiceCard from "../components/ServiceCard";
import { removeSessionKey } from "../utils/auth";
import RequestCard from "../components/RequestCard";
import { fetchWithAuth } from "../utils/authFetch";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/ReactToastify.css";
import { getAccountId } from "../utils/auth";

const HomePage = () => {
  const roomNum = sessionStorage.getItem("roomNum"); //room key of user
  //State for requests
  const [inputValue, setInputValue] = useState("");

  const RoomServiceTag = "Room Cleaning";
  const FoodDeliveryTag = "Food Delivery";
  const WakeUpCallTag = "Wake Up Call";
  const LaundryServiceTag = "Laundry Service";
  const SpaMassageTag = "Spa And Massage";
  const MaintenanceTag = "Maintenance";

  const [mainDish, setMainDish] = useState("");
  const [sideDish, setSideDish] = useState("");
  const [drink, setDrink] = useState("");
  const [mainChecked, setMainChecked] = useState(false);
  const [sideChecked, setSideChecked] = useState(false);
  const [drinkChecked, setDrinkChecked] = useState(false);

  // State for side drawers
  const [isMenuOpen, setMenuOpen] = useState(false);
  const [isProfileOpen, setProfileOpen] = useState(false);

  // Toggle side drawers
  const toggleMenu = () => setMenuOpen(!isMenuOpen);
  const toggleProfile = () => setProfileOpen(!isProfileOpen);

  // Close side drawers when clicking outside
  const closeMenu = () => setMenuOpen(false);
  const closeProfile = () => setProfileOpen(false);

  const roomInfo = getRoomInfo();

  const handleLogout = () => {
    removeSessionKey();
    window.location.href = "/";
  };

  const handleInputChange = (event) => {
    setInputValue(event.target.value);
  };

  const handleModalClose = () => {
    setInputValue("");
    setMainDish("");
    setSideDish("");
    setDrink("");

    setMainChecked(false);
    setSideChecked(false);
    setDrinkChecked(false);
  };

  const handleSubmit = async (tag) => {
    let items = "";

    if (tag === FoodDeliveryTag) {
      items = handleFoodDelivery();
    }
    const isInvalidRequest =
      (tag === WakeUpCallTag && !inputValue) ||
      (tag === FoodDeliveryTag && !items) ||
      (tag === MaintenanceTag && !inputValue);

    if (isInvalidRequest) {
      toast.error("Can't send your request: please enter a valid description!");
      setInputValue("");
      return;
    }

    const accountId = getAccountId(); // Retrieve accountId from the token
    if (!accountId && accountId !== 0) {
      toast.error("Unable to submit request. No account ID found.");
      return;
    }

    const requestBody = {
      taskType: tag.replace(/\s+/g, ""),
      description: items || inputValue || "N/A",
      roomId: parseInt(roomNum, 10),
      requesterId: accountId,
    };

    try {
      const response = await fetchWithAuth(
        `/tasks/`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(requestBody),
        }
      );

      if (response.ok) {
        toast.success("Successfully submitted request!");
        setInputValue("");
        resetFoodDeliveryInputs();
      } else {
        throw new Error("Failed to submit request");
      }
    } catch (error) {
      toast.error("Couldn't submit your request at this time!");
      console.error(error);
    }
  };

  const handleFoodDelivery = () => {
    let items = "";

    if (mainChecked && mainDish) {
      items += `Main: ${mainDish}. `;
    }
    if (sideChecked && sideDish) {
      items += `Side: ${sideDish}. `;
    }
    if (drinkChecked && drink) {
      items += `Drink: ${drink}. `;
    }

    return items.trim();
  };

  const resetFoodDeliveryInputs = () => {
    setMainDish("");
    setSideDish("");
    setDrink("");
    setMainChecked(false);
    setSideChecked(false);
    setDrinkChecked(false);
  };

  return (
    <div className="h-screen relative flex flex-col">
      <ToastContainer bodyClassName="toast-body" />
      <div className="flex-grow overflow-y-auto">
        {/* Sticky Header */}
        <header className="bg-white sticky top-0 my-2 p-2 flex justify-between items-center z-40">
          <button
            className="p-2 h-10 w-10 flex items-center justify-center"
            onClick={toggleMenu}
          >
            <FontAwesomeIcon icon={faBars} className="text-xl" />
          </button>
          <h1 className="text-2xl font-extrabold">Quick Service</h1>
          <button
            onClick={toggleProfile}
            className="rounded-full bg-gray-300 p-2 h-10 w-10 flex items-center justify-center"
          >
            <FontAwesomeIcon icon={faUser} className="text-xl text-white" />
          </button>
        </header>

        {/* Light Brown Section */}
        <div className="px-4 pt-3 pb-2 bg-primary rounded-t-[50%] mx-auto max-w-full shadow-[0_-2px_4px_rgba(0,0,0,0.2)] z-10">
          <div className="text-brown text-2xl font-bold text-center">
            {roomInfo.roomNumber}
          </div>
          <div className="text-gray-600 text-center text-sm">
            <FontAwesomeIcon icon={faCalendar} className="mr-2" />
            {roomInfo.periodOfStay}
          </div>
          <h2 className="mt-2 text-lg lg:text-xl font-semibold lg:text-center">
            Explore Our Services
          </h2>
          <p className="text-sm lg:text-md lg:text-center">
            Choose your service. We will deliver right to your door!
          </p>
        </div>

        <div className="absolute bottom-0 w-full h-1/2 bg-primary z-0"></div>

        {/* Requests Cards */}
        <div className="bg-primary">
          <div className="grid grid-cols-2 sm:grid-cols-3 gap-3 px-5 py-2 lg:py-5 mx-auto justify-items-center max-w-full lg:max-w-[50%] z-10">
            <RequestCard
              icon={faBroom}
              text={RoomServiceTag}
              onSubmit={handleSubmit}
              onClose={handleModalClose}
            >
              <label htmlFor="request" className="block mb-2">
                Special Instructions:
              </label>
              <textarea
                placeholder="Enter special instructions"
                value={inputValue}
                onChange={handleInputChange}
                className="border rounded p-2 mb-4 w-full h-[150px] resize-none"
              />
            </RequestCard>

            <RequestCard
              icon={faHamburger}
              text={FoodDeliveryTag}
              onSubmit={handleSubmit}
              onClose={handleModalClose}
            >
              <label htmlFor="request" className="block mb-2">
                Choose from our selection:
              </label>

              {/* Main Dish Selection */}
              <div className="flex items-center mb-4">
                <select
                  value={mainDish}
                  onChange={(e) => setMainDish(e.target.value)}
                  className="border rounded p-2 mr-2 w-48"
                  disabled={!mainChecked} // Disabled unless checkbox is checked
                >
                  <option value="" disabled>
                    Select a Main Dish
                  </option>
                  <option value="Grilled Chicken">Grilled Chicken</option>
                  <option value="Steak">Steak</option>
                  <option value="Pasta Primavera">Pasta Primavera</option>
                  <option value="Grilled Salmon">Grilled Salmon</option>
                </select>

                <label className="flex items-center">
                  <input
                    type="checkbox"
                    checked={mainChecked}
                    onChange={() => {
                      setMainChecked(!mainChecked);
                      if (mainChecked) setMainDish(""); // Reset dropdown when unchecked
                    }}
                    className="mr-2 w-5 h-5"
                  />
                  <span>Main</span>
                </label>
              </div>

              {/* Side Dish Selection */}
              <div className="flex items-center mb-4">
                <select
                  value={sideDish}
                  onChange={(e) => setSideDish(e.target.value)}
                  className="border rounded p-2 mr-2 w-48"
                  disabled={!sideChecked} // Disabled unless checkbox is checked
                >
                  <option value="" disabled>
                    Select a Side Dish
                  </option>
                  <option value="French Fries">French Fries</option>
                  <option value="Caesar Salad">Caesar Salad</option>
                  <option value="Steamed Vegetables">Steamed Vegetables</option>
                  <option value="Garlic Rice">Garlic Rice</option>
                </select>

                <label className="flex items-center">
                  <input
                    type="checkbox"
                    checked={sideChecked}
                    onChange={() => {
                      setSideChecked(!sideChecked);
                      if (sideChecked) setSideDish(""); // Reset dropdown when unchecked
                    }}
                    className="mr-2 w-5 h-5"
                  />
                  <span>Side</span>
                </label>
              </div>

              {/* Drink Selection */}
              <div className="flex items-center mb-4">
                <select
                  value={drink}
                  onChange={(e) => setDrink(e.target.value)}
                  className="border rounded p-2 mr-2 w-48"
                  disabled={!drinkChecked} // Disabled unless checkbox is checked
                >
                  <option value="" disabled>
                    Select a Drink
                  </option>
                  <option value="Soda">Soda</option>
                  <option value="Red Wine">Red Wine</option>
                  <option value="Cocktail">Cocktail</option>
                  <option value="Sparkling Water">Sparkling Water</option>
                </select>

                <label className="flex items-center">
                  <input
                    type="checkbox"
                    checked={drinkChecked}
                    onChange={() => {
                      setDrinkChecked(!drinkChecked);
                      if (drinkChecked) setDrink(""); // Reset dropdown when unchecked
                    }}
                    className="mr-2 w-5 h-5"
                  />
                  <span>Drink</span>
                </label>
              </div>
            </RequestCard>

            <RequestCard
              icon={faClock}
              text={WakeUpCallTag}
              onSubmit={handleSubmit}
              onClose={handleModalClose}
            >
              <label htmlFor="request" className="block mb-2">
                What time do we wake you up?
              </label>
              <select
                className="border rounded p-2 mb-4 w-full"
                value={inputValue}
                onChange={handleInputChange}
              >
                <option value="">Choose a time</option>
                <option value="1:00 AM">1:00 AM</option>
                <option value="2:00 AM">2:00 AM</option>
                <option value="3:00 AM">3:00 AM</option>
                <option value="4:00 AM">4:00 AM</option>
                <option value="5:00 AM">5:00 AM</option>
                <option value="6:00 AM">6:00 AM</option>
                <option value="7:00 AM">7:00 AM</option>
                <option value="8:00 AM">8:00 AM</option>
                <option value="9:00 AM">9:00 AM</option>
                <option value="10:00 AM">10:00 AM</option>
              </select>
            </RequestCard>

            <RequestCard
              icon={faShirt}
              text={LaundryServiceTag}
              onSubmit={handleSubmit}
              onClose={handleModalClose}
            >
              <label htmlFor="request" className="block mb-2">
                Special Instructions:
              </label>
              <textarea
                placeholder="Enter special instructions"
                value={inputValue}
                onChange={handleInputChange}
                className="border rounded p-2 mb-4 w-full h-[150px] resize-none"
              />
            </RequestCard>

            <RequestCard
              icon={faSpa}
              text={SpaMassageTag}
              onSubmit={handleSubmit}
              onClose={handleModalClose}
            >
              <label htmlFor="request" className="block mb-2">
                Special Instructions:
              </label>
              <textarea
                placeholder="Enter special instructions"
                value={inputValue}
                onChange={handleInputChange}
                className="border rounded p-2 mb-4 w-full h-[150px] resize-none"
              />
            </RequestCard>

            <RequestCard
              icon={faWrench}
              text={MaintenanceTag}
              onSubmit={handleSubmit}
              onClose={handleModalClose}
            >
              <label htmlFor="request" className="block mb-2">
                Details:
              </label>
              <textarea
                placeholder="Enter special instructions"
                value={inputValue}
                onChange={handleInputChange}
                className="border rounded p-2 mb-4 w-full h-[150px] resize-none"
              />
            </RequestCard>
          </div>
        </div>
      </div>

      {/* Left Side Drawer (Menu) */}
      {isMenuOpen && (
        <div
          className="fixed inset-0 bg-gray-800 bg-opacity-50 z-50"
          onClick={closeMenu}
        >
          <div
            className="fixed left-5 top-14 bg-white z-60 p-5 lg:p-8 text-xl lg:text-2xl rounded-lg"
            onClick={(e) => e.stopPropagation()}
          >
            <ul>
              <li>
                <a
                  href="#"
                  onClick={(e) => {
                    e.preventDefault(); // Prevent default navigation
                    toast.info("Coming soon!");
                  }}
                >
                  Menu
                </a>
              </li>
            </ul>
          </div>
        </div>
      )}

      {/* Right Side Drawer (Profile) */}
      {isProfileOpen && (
        <div
          className="fixed inset-0 bg-gray-800 bg-opacity-50 z-50"
          onClick={closeProfile}
        >
          <div
            className="fixed right-5 top-14 bg-white z-60 p-5 lg:p-8 text-xl lg:text-2xl rounded-lg"
            onClick={(e) => e.stopPropagation()}
          >
            <ul className="flex flex-col space-y-4">
              <li>
                <FontAwesomeIcon icon={faCog} className="mr-2" />
                <a
                  href="#"
                  onClick={(e) => {
                    e.preventDefault(); // Prevent default navigation
                    toast.info("Coming soon!");
                  }}
                >
                  Settings
                </a>
              </li>
              <li>
                <FontAwesomeIcon icon={faSignOutAlt} className="mr-2" />
                <a href="#" onClick={handleLogout}>
                  Log Out
                </a>
              </li>
            </ul>
          </div>
        </div>
      )}

      {/* Sticky Footer */}
      <footer className="grid grid-cols-3 space-x-0.5 items-center z-40 px-1 pb-1 ">
        <ServiceCard
          icon={faConciergeBell}
          text="Amenities"
          link="/amenities"
        />
        <ServiceCard
          icon={faClipboardList}
          text="Your requests"
          link="/tasks"
        />
        <ServiceCard
          icon={faExclamationTriangle}
          text="Complaint"
          link="/incident_reports"
        />
      </footer>
    </div>
  );
};

const getRoomInfo = () => {
  // Semi hard code stay duration: Right now it gets today time and next 7days
  const today = new Date();

  function formatDate(date) {
    const options = {
      day: "2-digit",
      month: "2-digit",
      year: "numeric",
      hour: "2-digit",
      minute: "2-digit",
      hour12: true,
    };
    return date.toLocaleString("en-GB", options).toLowerCase().replace(",", "");
  }

  const sevenDaysLater = new Date(today);
  sevenDaysLater.setDate(today.getDate() + 7);

  const todayStr = formatDate(today);
  const next7DaysStr = formatDate(sevenDaysLater);

  const periodOfStay = `${todayStr} - ${next7DaysStr}`;
  return {
    roomNumber: "Room " + sessionStorage.getItem("roomNum"),
    periodOfStay: periodOfStay,
  };
};

export default HomePage;
