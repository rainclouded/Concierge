import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Link } from "react-router-dom";

const ServiceCard = ({ icon, text, link }) => (
  <Link
    to={link}
    className="tile-item flex flex-col justify-center items-center bg-black text-white p-4 rounded-2xl transform transition-transform duration-300 hover:scale-105 w-full min-h-[150px] sm:min-h-[200px]"
  >
    <FontAwesomeIcon icon={icon} size="2x" />
    <p className="mt-4 text-lg sm:text-3xl">{text}</p>
  </Link>
);

export default ServiceCard;
