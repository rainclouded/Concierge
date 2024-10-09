import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";

import LoginPage from "./pages/LoginPage.jsx";
import HomePage from "./pages/HomePage.jsx";
import AmenitiesPage from "./pages/AmenitiesPage.jsx";
import IncidentReportPage from "./pages/IncidentReportPage";

import { createBrowserRouter, RouterProvider } from "react-router-dom";

const router = createBrowserRouter([
  {
    path: "/",
    element: <LoginPage />,
  },
  {
    path: "/home",
    element: <HomePage />,
  },
  {
    path: "/amenities",
    element: <AmenitiesPage />,
  },
  ,
  {
    path: "/incident_reports",
    element: <IncidentReportPage />,
  },
]);

createRoot(document.getElementById("root")).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>
);

