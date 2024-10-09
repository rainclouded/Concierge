import React from "react";
import AmenitiesPage from "../../src/pages/AmenitiesPage";
import { mount } from "cypress/react";
import { MemoryRouter } from "react-router-dom";

describe("AmenitiesPage Component Tests", () => {
  const mockAmenities = [
    {
      id: 1,
      name: "Pool",
      description: "Outdoor pool",
      startTime: "09:00:00",
      endTime: "21:00:00",
    },
    {
      id: 2,
      name: "Gym",
      description: "24/7 access gym",
      startTime: "00:00:00",
      endTime: "23:59:59",
    },
  ];

  beforeEach(() => {
    cy.intercept("GET", "**/amenities", {
      statusCode: 200,
      body: {
        message: "Amenities retrieved successfully.",
        data: mockAmenities,
      },
    }).as("getAmenities");

    // Wrap the AmenitiesPage component with MemoryRouter
    mount(
      <MemoryRouter>
        <AmenitiesPage />
      </MemoryRouter>
    );
  });

  it("should display the page title", () => {
    cy.get("h1").contains("Amenities");
  });

  it("should load the amenities table with correct headers", () => {
    // Check table headers
    cy.get("th").should(($th) => {
      expect($th).to.have.length(4); // 4 columns: Name, Description, Start Time, End Time
      expect($th.eq(0)).to.contain("Name");
      expect($th.eq(1)).to.contain("Description");
      expect($th.eq(2)).to.contain("Start Time");
      expect($th.eq(3)).to.contain("End Time");
    });
  });

  it("should display the correct number of rows based on mock data", () => {
    cy.wait("@getAmenities");
    cy.get("tbody tr").should("have.length", mockAmenities.length);
  });

  it('should display the "Back to Homepage" button', () => {
    cy.get("a").contains("Back to Homepage");
  });
});