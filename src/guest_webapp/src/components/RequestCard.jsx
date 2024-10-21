import React, { useState } from "react";
import Modal from "./Modal"; 
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

const RequestCard = ({text, icon, children }) => {
  const [isModalOpen, setModalOpen] = useState(false);

  const toggleModal = () => {
    setModalOpen(!isModalOpen);
  };

  return (
    <>
        <button 
            className="tile-item flex flex-col justify-center items-center bg-black text-white p-4 rounded-2xl transform transition-transform duration-300 hover:scale-105 w-full min-h-[150px] sm:min-h-[200px]"
            onClick={toggleModal}
            >
            <FontAwesomeIcon icon={icon} size="2x" />
            <p className="mt-4 text-lg sm:text-3xl">{text}</p>
        </button>

        <Modal isOpen={isModalOpen} onClose={toggleModal}>
        <h2 className="text-xl font-bold mb-4">Request {text}</h2>
            {children}
        <button
            className="bg-black text-white p-4 rounded-2xl mt-4 text-small sm:text-1xl"
          >
            Submit Request
          </button> 
        </Modal>
    </>
  );
};

export default RequestCard;
