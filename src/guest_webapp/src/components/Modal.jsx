import React from "react";

const Modal = ({ isOpen, onClose, children }) => {
  if (!isOpen) return null; 

  return (
    <div className="fixed inset-0 bg-gray-800 bg-opacity-50 z-50" onClick={onClose}>
      <div
        className="fixed left-1/2 top-1/2 transform -translate-x-1/2 -translate-y-1/2 bg-white rounded-lg shadow-lg z-60 aspect-square w-[95%] lg:w-[400px]"
        onClick={(e) => e.stopPropagation()}
      >
        <button className="absolute top-2 right-4 text-3xl font-medium text-red-700" onClick={onClose}>
          &times; 
        </button>
        <div className="p-6 flex flex-col h-full">
          {children}
        </div>
      </div>
    </div>
  );
};

export default Modal;