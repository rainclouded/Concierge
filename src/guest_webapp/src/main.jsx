import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";

import LoginPage from "./pages/LoginPage.jsx";
import HomePage from "./pages/HomePage.jsx";
import AmenitiesPage from "./pages/AmenitiesPage.jsx";
import IncidentReportPage from "./pages/IncidentReportPage";

import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { ProtectedRoute } from "./utils/ProtectedRoutes.jsx";
import { isAuthenticated } from "./utils/auth.js";

const router = createBrowserRouter([
  {
    path: "/",
    element: <ProtectedRoute canAccessFunc={()=>!isAuthenticated()} element={<LoginPage />} fallbackPath={"/home"} />,
  },
  {
    path: "/home",
    element: <ProtectedRoute canAccessFunc={isAuthenticated} element={<HomePage />} fallbackPath={"/"} />,
  },
  {
    path: "/amenities",
    element: <ProtectedRoute canAccessFunc={isAuthenticated} element={<AmenitiesPage />} fallbackPath={"/"} />,
  },
  {
    path: "/incident_reports",
    element: <ProtectedRoute canAccessFunc={isAuthenticated} element={<IncidentReportPage />} fallbackPath={"/"} />,
  },
]);

createRoot(document.getElementById("root")).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>
);
