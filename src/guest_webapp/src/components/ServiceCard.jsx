import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Link } from "react-router-dom";

const ServiceCard = ({ icon, text, link }) => (
  <Link
    to={link}
    className="tile-item flex flex-col lg:flex-row justify-center items-center text-center border bg-secondary text-black p-4 rounded-md transform transition-transform duration-300"
  >
    <FontAwesomeIcon icon={icon} className="text-3xl" />
    <p className="text-sm lg:ml-2 lg:text-xl">{text}</p>
  </Link>
);

export default ServiceCard;
