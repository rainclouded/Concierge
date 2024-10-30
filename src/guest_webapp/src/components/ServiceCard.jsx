import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Link } from "react-router-dom";

const ServiceCard = ({ icon, text, link }) => (
  <Link
    to={link}
    className="tile-item flex flex-col justify-center items-center bg-black text-white p-4 rounded-2xl transform transition-transform duration-300 hover:scale-105 min-w-[150px]"
  >
    <FontAwesomeIcon icon={icon} size="1x" />
    <p className="mt-1 text-small sm:text-1xl">{text}</p>
  </Link>
);

export default ServiceCard;
