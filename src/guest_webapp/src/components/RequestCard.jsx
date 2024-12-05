import React, { useState } from "react";
import Modal from "./Modal"; 
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

const RequestCard = ({ text, icon, children, onSubmit, onClose }) => {
  const [isModalOpen, setModalOpen] = useState(false);

  const toggleModal = () => {
    setModalOpen(!isModalOpen);
  };

  return (
    <>
      <button 
        className="tile-item flex flex-col justify-center items-center bg-black text-white p-4 rounded-2xl transform transition-transform duration-300 hover:scale-[102%] w-full aspect-square"
        onClick={toggleModal}
      >
        <FontAwesomeIcon icon={icon} size="2x" />
        <p className="mt-4 text-lg sm:text-3xl">{text}</p>
      </button>

      <Modal isOpen={isModalOpen} onClose={() => { toggleModal(); onClose(); }}>
        <h2 className="text-xl font-bold mb-4">Request {text}</h2>
        <div className="flex-grow h-full">
          {children}
        </div>
        <div className="mt-4"> 
          <button
            className="bg-black text-white p-4 rounded-2xl text-small sm:text-1xl w-full" 
            onClick={() => {
              onSubmit(text);
              toggleModal();
            }}
          >
            Submit Request
          </button> 
        </div>
      </Modal>
    </>
  );
};

export default RequestCard;