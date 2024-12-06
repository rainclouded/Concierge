describe("HomePage Integration Tests", () => {
  beforeEach(() => {
    cy.viewport(390, 844);
    cy.visit("localhost:8081");
    cy.get("#roomNum").type("11111");
    cy.get("#roomKey").type("password");
    cy.get("button").click();
    cy.url().should("include", "/home");
    cy.get(".tile-item").contains("Complaint").click();
  });

  it("POST incident repot", () => {
    cy.contains("Make a report");

    cy.get("form").find("input").clear().type("TestingIncidentCreation");
    cy.get("form")
      .find("textarea")
      .clear()
      .type("TestingIncidentCreationDescription");
    cy.contains("button", "Submit").click();
    // close toast since it blocks the button
    cy.get('.Toastify__close-button').should('be.visible').click();
  });

  afterEach(() => {
    cy.get(".svg-inline--fa.fa-arrow-left").click();
    cy.get("header").find(".rounded-full.bg-gray-300").click();
    cy.get('.fixed.right-5.top-14.bg-white').should('be.visible');
    cy.contains('Log Out').click();
    cy.url().should('include', '');
  });
});
