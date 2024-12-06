describe("HomePage Integration Tests", () => {
  beforeEach(() => {
    cy.viewport(390, 844);
    cy.visit("localhost:8081");
    cy.get("#roomNum").type("11111");
    cy.get("#roomKey").type("password");
    cy.get("button").click();
    cy.url().should("include", "/home");
    cy.get(".tile-item").contains("Your requests").click();
  });

  it("POST incident repot", () => {
    cy.contains("Your requests");

    cy.get(".flex.flex-col.space-y-4").should("exist");

    // Check if any task has the title "Room Cleaning"
    cy.get(".relative.min-h-36.bg-lightPrimary").within(() => {
      cy.contains(
        ".text-xl.font-semibold.tracking-wide",
        "Room Cleaning"
      ).should("exist");
    });
  });

  afterEach(() => {
    cy.get(".svg-inline--fa.fa-arrow-left").click();
    cy.get("header").find(".rounded-full.bg-gray-300").click();
    cy.get(".fixed.right-5.top-14.bg-white").should("be.visible");
    cy.contains("Log Out").click();
    cy.url().should("include", "");
  });
});
