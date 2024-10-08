import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import LoginPage from "./pages/LoginPage";
import HomePage from "./pages/HomePage";
import AmenitiesPage from "./pages/AmenitiesPage";
import IncidentReportPage from "./pages/IncidentReportPage";
import "./App.css";

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LoginPage />} />
        <Route path="/home" element={<HomePage />} />
        <Route path="/amenities" element={<AmenitiesPage />} />
        <Route path="/incident_reports" element={<IncidentReportPage />} />
      </Routes>
    </Router>
  );
};

export default App;
