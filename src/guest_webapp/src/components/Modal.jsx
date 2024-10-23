import React from "react";

const Modal = ({ isOpen, onClose, children }) => {
  if (!isOpen) return null; 

  return (
    <div className="fixed inset-0 bg-gray-800 bg-opacity-50 z-50" onClick={onClose}>
      <div
        className="fixed left-1/2 top-1/2 transform -translate-x-1/2 -translate-y-1/2 bg-white rounded-lg shadow-lg z-60"
        style={{ width: '400px', height: '400px', overflowY: 'auto' }}
        onClick={(e) => e.stopPropagation()}
      >
        <button className="absolute top-2 right-2" onClick={onClose}>
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