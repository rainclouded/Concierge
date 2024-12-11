describe("Home Page Integration Tests", () => {
  beforeEach(() => {
    cy.viewport(390, 844);
    cy.visit("localhost:8089");
    cy.get("#roomNum").type("11111");
    cy.get("#roomKey").type("password");
    cy.get("button").click();
    cy.url().should("include", "/home");
  });

  it("Get Quick Service Buttons", () => {
    cy.url().should("include", "home");
    cy.contains("button", "Room Cleaning").should("be.visible");
    cy.contains("button", "Food Delivery").should("be.visible");
    cy.contains("button", "Wake Up Call").should("be.visible");
    cy.contains("button", "Laundry Service").scrollIntoView();
    cy.contains("button", "Laundry Service").should("be.visible");
    cy.contains("button", "Spa And Massage").should("be.visible");
    cy.contains("button", "Maintenance").should("be.visible");
  });

  it("Submit request buttons", () => {
    cy.contains("button", "Room Cleaning").click();
    cy.contains("button", "Submit Request").click();
    // close toast since it blocks the button
    cy.get(".Toastify__close-button").should("be.visible").click();
  });

  it("Submit Wake Up Call request", () => {
    cy.contains("button", "Wake Up Call").click();
    cy.get('select').select('7:00 AM');
    cy.contains("button", "Submit Request").click();
    cy.get(".Toastify__close-button").should("be.visible").click();
  });

  it("Submit Food Delivery request", () => {
    cy.contains("button", "Food Delivery").click();
    cy.get('input[type="checkbox"]').first().check();
    cy.get('select').eq(0).select('Grilled Chicken');
    cy.get('input[type="checkbox"]').eq(1).check();
    cy.get('select').eq(1).select('French Fries');
    cy.get('input[type="checkbox"]').eq(2).check();
    cy.get('select').eq(2).select('Soda'); 
    cy.contains("button", "Submit Request").click();
    cy.get(".Toastify__close-button").should("be.visible").click();
  });
  
  afterEach(() => {
    cy.get("header").find(".rounded-full.bg-gray-300").click();
    cy.get(".fixed.right-5.top-14.bg-white").should("be.visible");
    cy.contains("Log Out").click();
    cy.url().should("include", "");
  });
});
