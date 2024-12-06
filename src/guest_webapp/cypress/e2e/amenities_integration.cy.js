describe("Amenities Integration Tests", () => {
  beforeEach(() => {
    cy.viewport(390, 844);
    cy.visit("localhost:8081");
    cy.get("#roomNum").type("11111");
    cy.get("#roomKey").type("password");
    cy.get("button").click();
    cy.url().should("include", "/home");
    cy.get(".tile-item").contains("Amenities").click();
  });

  it("GET amenities", () => {
    cy.contains("Amenities"); // Ensure the title is rendered

    // Assert the presence of amenity items
    cy.get(".rounded-lg.p-4.shadow-md").should(($items) => {
      expect($items).to.have.length(4); // Check that 4 amenities are displayed

      // Verify the text content of each amenity
      const amenities = ["Pool", "Gym", "Breakfast", "Bar"];
      $items.each((index, item) => {
        expect(item).to.contain.text(amenities[index]);
      });
    });
  });

  afterEach(() => {
    cy.get(".svg-inline--fa.fa-arrow-left").click();
    cy.get("header").find(".rounded-full.bg-gray-300").click();
    cy.get('.fixed.right-5.top-14.bg-white').should('be.visible');
    cy.contains('Log Out').click();
    cy.url().should('include', '');
  });
});
