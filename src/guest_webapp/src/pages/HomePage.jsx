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
  //State for requests
  const [inputValue, setInputValue] = useState('');

  const RoomServiceTag = "Room Service"
  const FoodDeliveryTag = "Food Delivery"
  const WakeUpCallTag = "Wake Up Call"
  const LaundryServiceTag = "Laundry Service"
  const SpaMassageTag = "Spa & Massage"
  const Maintenance = "Maintenance"

  const [mainDish, setMainDish] = useState('');
  const [sideDish, setSideDish] = useState('');
  const [drink, setDrink] = useState('');
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
    let items = '';
    if(tag === FoodDeliveryTag){

      console.log(mainDish)
        if (mainChecked && mainDish) {
          items += `Main: ${mainDish}. `;
        }
        if (sideChecked && sideDish) {
          items += `Side: ${sideDish}. `;
        }
        if (drinkChecked && drink) {
          items += `Drink: ${drink}. `;
        }
    }

    if((tag === WakeUpCallTag && inputValue === "") || (tag === FoodDeliveryTag && items === "")){
      alert("Can't send your request: please enter values!");
      setInputValue("");
      return;  
    }

    const requestBody = {
      title: tag, 
      description: tag === FoodDeliveryTag ? items.trim() : inputValue,
      room_id: roomInfo.roomNumber,
      requester_account_id: 100,
      status: "OPEN"
    };
    
    setInputValue("");

    try {
      const response = await fetch(
        `${import.meta.env.VITE_API_BASE_URL}/tasks/`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(requestBody),
        }
      );
      if (!response.ok) throw new Error("Failed to submit request")
    } catch (error) {
      alert("Couldn't submit your request at this time! ");
      console.log(error);
    }
  };

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
                  checked={mainChecked}
                  onChange={() => setMainChecked(!mainChecked)}
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
              >
                <option value="">Select a Side Dish</option>
                <option value="French Fries">French Fries</option>
                <option value="Caesar Salad">Caesar Salad</option>
                <option value="Steamed Vegetables">Steamed Vegetables</option>
                <option value="Garlic Rice">Garlic Rice</option>
              </select>

              <label className="flex items-center">
                <input
                  type="checkbox"
                  checked={sideChecked}
                  onChange={() => setSideChecked(!sideChecked)}
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
              >
                <option value="">Select a Drink</option>
                <option value="Soda">Soda</option>
                <option value="Red Wine">Red Wine</option>
                <option value="Cocktail">Cocktail</option>
                <option value="Sparkling Water">Sparkling Water</option>
              </select>

              <label className="flex items-center">
                <input
                  type="checkbox"
                  checked={drinkChecked}
                  onChange={() => setDrinkChecked(!drinkChecked)}
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
            <input
              placeholder="Enter special instructions"
              value={inputValue}
              onChange={handleInputChange}
              className="border rounded p-2 mb-4 w-full"
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
            <input
              placeholder="Enter special instructions"
              value={inputValue}
              onChange={handleInputChange}
              className="border rounded p-2 mb-4 w-full"
            />
          </RequestCard>   

          <RequestCard
            icon={faWrench}
            text={Maintenance}
            onSubmit={handleSubmit}
            onClose={handleModalClose}
          >
            <label htmlFor="request" className="block mb-2">
              Details:
            </label>
            <input
              placeholder="What needs to be fixed?"
              value={inputValue}
              onChange={handleInputChange}
              className="border rounded p-2 mb-4 w-full"
            />
          </RequestCard>    
        </div>
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

      {/* Sticky Footer */}
      <footer className="bg-white p-4 shadow-md flex justify-between items-center z-40">
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
      </footer>

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
