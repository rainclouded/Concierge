import HomePage from "../../src/pages/HomePage";
import { MemoryRouter } from "react-router-dom";

describe("HomePage Component Tests", () => {
  it("should render the HomePage component and display Quick Service", () => {
    // Wrap the HomePage with MemoryRouter to provide routing context
    cy.mount(
      <MemoryRouter>
        <HomePage />
      </MemoryRouter>
    );

    // Check if the "Quick Service" text is visible
    cy.contains("Quick Service").should("be.visible");
  });

  it("should display the room number and period of stay", () => {
    // Wrap the HomePage with MemoryRouter to provide routing context
    cy.mount(
      <MemoryRouter>
        <HomePage />
      </MemoryRouter>
    );

    // Check if the room number is visible
    cy.contains("Room 404").should("be.visible");

    // Check if the period of stay is visible
    cy.contains("23.11 11:00am - 26.11 11:00am").should("be.visible");
  });

  it('should render the "Amenities" and "Incident Report" service cards', () => {
    // Wrap the HomePage with MemoryRouter to provide routing context
    cy.mount(
      <MemoryRouter>
        <HomePage />
      </MemoryRouter>
    );

    // Check if "Amenities" card is visible
    cy.contains("Amenities").should("be.visible");

    // Check if "Incident Report" card is visible
    cy.contains("Incident Report").should("be.visible");
  });
});