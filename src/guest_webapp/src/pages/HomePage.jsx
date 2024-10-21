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
  faWrench
} from "@fortawesome/free-solid-svg-icons";

import ServiceCard from "../components/ServiceCard";
import RequestCard from "../components/RequestCard";

const HomePage = () => {
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

  return (
    <div className="min-h-screen bg-[#ECD8C8] relative">
      {/* Sticky Header */}
      <header className="sticky top-0 bg-white p-4 shadow-md flex justify-between items-center z-40">
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
          <FontAwesomeIcon icon={faUser} className="text-xl" />
        </button>
      </header>

      {/* Light Brown Section */}
      <div className="p-4 bg-[#ECD8C8] rounded-t-xl text-center mx-auto max-w-full md:max-w-[75%]">
        <div className="text-[#8F613C] text-2xl font-bold">
          {roomInfo.roomNumber}
        </div>
        <div className="text-gray-500">{roomInfo.periodOfStay}</div>
        <h2 className="mt-4 text-lg font-semibold">Explore Our Services</h2>
        <p>Choose your service. We will deliver right to your door!</p>
      </div>

      {/* Service Cards */}
      <div className="grid grid-cols-2 sm:grid-cols-3 gap-4 p-4 mx-auto justify-items-center max-w-full md:max-w-[75%]">
        <RequestCard
          icon={faBroom}
          text="Room Service"
        >
          <label htmlFor="roomNumber" className="block mb-2">
            Special Instructions:
          </label>
          <input
            type="text"
            className="border rounded p-2 mb-4 w-full"
            placeholder="Your instructions here..."
          />
        </RequestCard>          
        <RequestCard
          icon={faHamburger}
          text="Food Delivery"
        >
          <label htmlFor="roomNumber" className="block mb-2">
            Choose from our selection:
          </label>
          {/* Main Dish Selection */}
          <div className="flex items-center mb-4">
            <select className="border rounded p-2 mr-2 w-48"> {/* Fixed width */}
              <option value="">Select a Main Dish</option>
              <option value="grilledChicken">Grilled Chicken</option>
              <option value="steak">Steak</option>
              <option value="pasta">Pasta Primavera</option>
              <option value="salmon">Grilled Salmon</option>
            </select>

            <label className="flex items-center">
              <input
                type="checkbox"
                className="mr-2 w-5 h-5"
              />
              <span>Main</span>
            </label>
          </div>

          {/* Side Dish Selection */}
          <div className="flex items-center mb-4">
            <select className="border rounded p-2 mr-2 w-48"> {/* Fixed width */}
              <option value="">Select a Side Dish</option>
              <option value="fries">French Fries</option>
              <option value="salad">Caesar Salad</option>
              <option value="vegetables">Steamed Vegetables</option>
              <option value="rice">Garlic Rice</option>
            </select>

            <label className="flex items-center">
              <input
                type="checkbox"
                className="mr-2 w-5 h-5"
              />
              <span>Side</span>
            </label>
          </div>

          {/* Drink Selection */}
          <div className="flex items-center mb-4">
            <select className="border rounded p-2 mr-2 w-48"> {/* Fixed width */}
              <option value="">Select a Drink</option>
              <option value="soda">Soda</option>
              <option value="wine">Red Wine</option>
              <option value="cocktail">Cocktail</option>
              <option value="water">Sparkling Water</option>
            </select>

            <label className="flex items-center">
              <input
                type="checkbox"
                className="mr-2 w-5 h-5"
              />
              <span>Drink</span>
            </label>
          </div>
        </RequestCard>
        <RequestCard
          icon={faClock}
          text="Wake Up Call"
        >
          <label htmlFor="roomNumber" className="block mb-2">
            What time do we wake you up?
          </label>
          <select
            className="border rounded p-2 mb-4 w-full"
          >
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
          text="Laundry Service"
        >
          <label htmlFor="roomNumber" className="block mb-2">
            Special Instructions:
          </label>
          <input
            type="text"
            className="border rounded p-2 mb-4 w-full"
            placeholder="Your instructions here..."
          />
        </RequestCard>   
        <RequestCard
          icon={faSpa}
          text="Spa & Massage"
        >
          <label htmlFor="roomNumber" className="block mb-2">
            Special Instructions:
          </label>
          <input
            type="text"
            className="border rounded p-2 mb-4 w-full"
            placeholder="Your instructions here..."
          />
        </RequestCard>   
        <RequestCard
          icon={faWrench}
          text="Maintenance"
        >
          <label htmlFor="roomNumber" className="block mb-2">
            Special Instructions:
          </label>
          <input
            type="text"
            className="border rounded p-2 mb-4 w-full"
            placeholder="Your instructions here..."
          />
        </RequestCard>   
        <ServiceCard
          icon={faConciergeBell}
          text="Amenities"
          link="/amenities"
        />
        <ServiceCard
          icon={faExclamationTriangle}
          text="Report an Incident"
          link="/incident_reports"
        />
      </div>
      
      {/* Left Side Drawer (Menu) */}
      {isMenuOpen && (
        <div
          className="fixed inset-0 bg-gray-800 bg-opacity-50 z-50"
          onClick={closeMenu}
        >
          <div
            className="fixed left-0 top-0 h-full w-1/2 md:w-1/4 bg-white z-60 p-4"
            onClick={(e) => e.stopPropagation()}
          >
            <ul>
              <li>
                <a href="/home">Menu</a>
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
            className="fixed right-0 top-0 h-full w-3/4 md:w-1/4 bg-white z-60 p-4"
            onClick={(e) => e.stopPropagation()}
          >
            <ul>
              <li>
                <a href="/home">Settings</a>
              </li>
              <li>
                <a href="/">Log Out</a>
              </li>
            </ul>
          </div>
        </div>
      )}
    </div>
  );
  
};

// hardcoded, change later
const getRoomInfo = () => {
  return {
    roomNumber: "Room 404",
    periodOfStay: "23.11  11:00am - 26.11  11:00am",
  };
};

export default HomePage;
